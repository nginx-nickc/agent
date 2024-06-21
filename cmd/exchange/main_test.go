package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"testing"

	"google.golang.org/grpc"

	grpc2 "github.com/nginx/agent/v3/internal/grpc"
)

func BenchmarkStreamUpload(t *testing.B) {
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

	for i := 0; i < t.N; i++ {
		_ = streamUpload(ctx, conn)
	}
}

func BenchmarkRPCUpload(t *testing.B) {
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

	for i := 0; i < t.N; i++ {
		_ = rpcUpload(ctx, conn)
	}
}
