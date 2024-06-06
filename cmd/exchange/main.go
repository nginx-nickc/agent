/*
 *  Copyright 2024 F5, Inc. All rights reserved.
 *
 *  No part of the software may be reproduced or transmitted in any
 *  form or by any means, electronic or mechanical, for any purpose,
 *  without express written permission of F5 Networks, Inc.
 *
 */

package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/nginx/agent/v3/api/grpc/mpi/v1"
	grpc2 "github.com/nginx/agent/v3/internal/grpc"
	"github.com/nginx/agent/v3/pkg/collections"
	"github.com/nginx/agent/v3/pkg/grpc/messenger"
)

//go:embed testfiles/*
var configs embed.FS

const (
	token      = ""
	instanceID = "8525cd7a-f221-42f6-ac52-4d97a6a169fa"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	conn, err := grpc.NewClient("localhost:8888",
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(
			&grpc2.PerRPCCredentials{
				Token: token,
				ID:    instanceID,
			}),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//_ = readFiles()
	var rootCmd = &cobra.Command{
		Use:   "exchange",
		Short: "Exchange is a gRPC test client for connecting to local N1SC dataplane-ctrl",
		Long:  `Exchange is a gRPC test client for connecting to local N1SC dataplane-ctrl`,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	var streamStandAlone = &cobra.Command{
		Use:   "stream-standalone",
		Short: "Standalone stream",
		Long:  `Standalone stream, where the upload/download is independent`,
		Run: func(cmd *cobra.Command, args []string) {
			op := "upload"
			if len(args) > 0 && args[0] == "download" {
				op = "download"
			}
			switch op {
			case "upload":
				if err = streamStandaloneUpload(ctx, conn); err != nil {
					log.Fatal(err)
				}
			case "download":
				if err = streamStandaloneDownload(ctx, conn); err != nil {
					log.Fatal(err)
				}
			}
		},
	}
	rootCmd.AddCommand(streamStandAlone)
	if err = rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
func streamStandaloneUpload(parent context.Context, conn *grpc.ClientConn) error {
	ctx, cancel := context.WithCancel(parent)
	defer cancel()

	client := pb.NewCommandServiceClient(conn)
	exchangeClient, err := client.FileExchange(ctx)
	if err != nil {
		log.Fatal(err)
	}
	me := messenger.New[pb.Exchange, pb.Exchange](5, exchangeClient)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go me.Run(ctx, &wg)

	fm := readFiles()
	fo := make([]*pb.File, 0, len(fm))
	for _, fr := range fm {
		FMeta := fr.Response.File.Meta
		fo = append(fo, &pb.File{
			FileMeta: FMeta,
		})
	}
	if err = me.Send(ctx, &pb.Exchange{
		Meta: &pb.MessageMeta{
			MessageId:     uuid.NewString(),
			CorrelationId: uuid.NewString(),
			Timestamp:     timestamppb.New(time.Now()),
		},
		Message: &pb.Exchange_Overview{
			Overview: &pb.FileOverview{
				Files: fo,
			},
		},
	}); err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err = <-me.Errors():
			log.Fatalf("Received error: %s\n", err)
		case msg := <-me.Messages():
			switch m := msg.Msg.Message.(type) {
			case *pb.Exchange_Request:
				for _, fr := range m.Request.GetFiles() {
					ffr, ok := fm[fr.Name]
					if !ok {
						return status.Error(codes.NotFound, fmt.Sprintf("%s not found", fr.Name))
					}
					ff := ffr.Response.File
					err = me.Send(ctx, &pb.Exchange{
						Meta: msg.Msg.Meta,
						Message: &pb.Exchange_Response{
							Response: &pb.FileResponse{
								File: &pb.FileResponse_File{
									Meta: &pb.FileMeta{
										Name:         fr.Name,
										Hash:         ff.Meta.Hash,
										ModifiedTime: ff.Meta.ModifiedTime,
										Size:         int64(len(ff.File.Contents)),
									},
									File: &pb.FileContents{Contents: ff.File.Contents},
								},
							},
						},
					})
				}
				if err != nil {
					return status.Error(codes.Aborted, err.Error())
				}
			}
		}
	}
	return nil
}

func streamStandaloneDownload(parent context.Context, conn *grpc.ClientConn) error {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(parent)
	defer func() {
		cancel()
		wg.Wait()
	}()

	client := pb.NewCommandServiceClient(conn)
	exchangeClient, err := client.FileExchange(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File exchange stream started", conn.GetState())
	me := messenger.New[pb.Exchange, pb.Exchange](5, exchangeClient)
	wg.Add(1)
	go me.Run(ctx, &wg)
	if err = me.Send(ctx, &pb.Exchange{
		Meta: &pb.MessageMeta{
			MessageId:     uuid.NewString(),
			CorrelationId: uuid.NewString(),
			Timestamp:     timestamppb.New(time.Now()),
		},
		Message: &pb.Exchange_OverviewReq{},
	}); err != nil {
		log.Fatal(err)
	}
	var reqQuestMap map[string]struct{}
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err = <-me.Errors():
			log.Fatalf("Received error: %s\n", err)
		case msg := <-me.Messages():
			log.Printf("Received message: %T\n", msg)
			switch m := msg.Msg.Message.(type) {
			case *pb.Exchange_Overview:
				// we got an overview from the management server, we should mimic retrieval
				files := make([]*pb.FileMeta, 0, len(m.Overview.Files))
				reqQuestMap = make(map[string]struct{})
				for _, f := range m.Overview.Files {
					files = append(files, f.FileMeta)
					reqQuestMap[f.FileMeta.Name] = struct{}{}
				}
				if len(reqQuestMap) == 0 {
					log.Println("no files exchanged needed")
					return nil
				}
				req := &pb.Exchange_Request{Request: &pb.FilesRequest{Files: files}}
				if err = me.Send(ctx, &pb.Exchange{
					Meta: &pb.MessageMeta{
						MessageId:     uuid.NewString(),
						CorrelationId: uuid.NewString(),
						Timestamp:     timestamppb.New(time.Now()),
					},
					Message: req,
				}); err != nil {
					log.Fatal(err)
				}
			case *pb.Exchange_Response:
				f := m.Response.File
				fn := f.Meta.Name
				if _, ok := reqQuestMap[fn]; !ok {
					log.Printf("unexpected file response: %s\n", fn)
				}
				tmpPrefix := filepath.Join(os.TempDir(), uuid.NewString())
				tmpFN := filepath.Join(tmpPrefix, fn)
				if os.MkdirAll(filepath.Dir(tmpFN), 0777) != nil {
					log.Fatal(err)
				}
				log.Println("create directory:", tmpFN)
				var ff *os.File
				ff, err = os.Create(tmpFN)
				if err != nil {
					log.Fatal(err)
				}
				_, err = ff.Write(f.File.Contents)
				if err != nil {
					log.Fatal(err)
				}
				delete(reqQuestMap, fn)
				if len(reqQuestMap) == 0 {
					log.Println("done file transfer")
					return nil
				}
			}
		case <-time.Tick(30 * time.Second):
			fmt.Println("state", conn.GetState())
			err = me.Send(ctx, &pb.Exchange{
				Meta:    nil,
				Message: nil,
			})
			fmt.Println("state", conn.GetState(), err)
		}
	}
	return nil
}

//func download() {
//	log.Println("starting exchange")
//	client := pb.NewCommandServiceClient(conn)
//	exchangeClient, err := client.FileExchange(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//	me := messenger.New[pb.Exchange, pb.Exchange](5, exchangeClient)
//	wg := sync.WaitGroup{}
//	wg.Add(1)
//	go me.Run(ctx, &wg)
//
//	log.Println("starting sending")
//	for _, fn := range fm {
//		err = me.Send(ctx, &pb.Exchange{
//			Meta:    nil,
//			Message: fn,
//		})
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
//	wg.Wait()
//}

func readFiles() map[string]*pb.Exchange_Response {
	files := make(map[string]*pb.Exchange_Response)
	base := "testfiles"
	if err := fs.WalkDir(configs, base,
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			content, err := configs.ReadFile(path)
			if err != nil {
				return err
			}
			stats, err := os.Stat(path)
			if err != nil {
				return err
			}
			fn, err := filepath.Rel(base, path)
			if err != nil {
				return err
			}
			fn = filepath.Join("/", fn)
			files[fn] = &pb.Exchange_Response{
				Response: &pb.FileResponse{
					File: &pb.FileResponse_File{
						Meta: &pb.FileMeta{
							Name:         fn,
							Hash:         collections.Hash(content),
							ModifiedTime: timestamppb.New(stats.ModTime()),
							Size:         int64(len(content)),
						},
						File: &pb.FileContents{Contents: content},
					},
				},
			}
			fmt.Println(path, fn)

			return nil

		}); err != nil {
		log.Fatal(err)
	}
	return files
}
