/*
 * Copyright 2023 F5, Inc. All rights reserved.
 *
 * No part of the software may be reproduced or transmitted in any
 * form or by any means, electronic or mechanical, for any purpose,
 * without express written permission of F5 Networks, Inc.
 */

// Package messenger provides generic message handling abstractions for gRPC streams.
package messenger

import (
	"context"
	"errors"
	"sync"
)

type Message[T any] struct {
	Msg *T
}

// Handler keep track of incoming/outgoing gRPC messages. The wrapper Message can be used by both message going to or
// coming from the handler, outgoing to the remote (nginx-agent) or incoming from the remote (nginx-agent) respectively.
type Handler[In, Out any] struct {
	incoming chan *Message[In]  // incoming to handler (-> handler) , send using the sender
	outgoing chan *Message[Out] // outgoing from handler ( <- handler ), receive from receiver
	error    chan error         // error from receive
	send     Sender[In]
	receiver Receiver[Out]
}

type Receiver[Out any] interface {
	Recv() (*Out, error)
}

type Sender[In any] interface {
	Send(*In) error
}

type SendReceiver[In, Out any] interface {
	Sender[In]
	Receiver[Out]
}

type SendReceiverInt interface {
	SendReceiver[int, int]
}

func New[In, Out any](buffSize int, sr SendReceiver[In, Out]) *Handler[In, Out] {
	return &Handler[In, Out]{
		incoming: make(chan *Message[In], buffSize),
		outgoing: make(chan *Message[Out], buffSize),
		error:    make(chan error),
		send:     sr,
		receiver: sr,
	}
}

// Run starts up the handlers. We listen on the receiver, and incoming chan. Assume we are running inside a goroutine.
// Presumably the SendReceiver is a message specific gRPC stream, like:
// - https://github.com/nginx/agent/blob/main/sdk/proto/command_svc.pb.go#L88
// The provided context should be from the underlying gRPC stream.
// For references:
// - https://pkg.go.dev/google.golang.org/grpc#ClientStream
// - https://pkg.go.dev/google.golang.org/grpc#ServerStream
//
// RecvMsg blocks until it receives a message into m or the stream is
// done. It returns io.EOF when the stream completes successfully. On
// any other error, the stream is aborted and the error contains the RPC
// status.
// SendMsg blocks until:
//   - There is sufficient flow control to schedule m with the transport, or
//   - The stream is done, or
//   - The stream breaks.
//
// SendMsg does not wait until the message is received by the server. An
// untimely stream closure may result in lost messages. To ensure delivery,
// users should ensure the RPC completed successfully using RecvMsg.
//
// It is safe to have a goroutine calling SendMsg and another goroutine
// calling RecvMsg on the same stream at the same time, but it is not safe
// to call SendMsg on the same stream in different goroutines. It is also
// not safe to call CloseSend concurrently with SendMsg.
func (h *Handler[In, Out]) Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	wg.Add(1)
	go h.handleRecv(ctx, wg)
	h.handleSend(ctx, wg)
}

// Send a message, will return error if the context is Done. Can block if the incoming chan is full, use context
// for control.
func (h *Handler[In, Out]) Send(ctx context.Context, msg *In) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case h.incoming <- &Message[In]{
		Msg: msg,
	}:
	}
	return nil
}

// handleSend handles outgoing to the sender.
func (h *Handler[In, Out]) handleSend(ctx context.Context, _ *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			return
		case m := <-h.incoming:
			err := h.send.Send(m.Msg)
			if err != nil {
				if errors.Is(err, context.Canceled) || errors.Is(ctx.Err(), context.Canceled) {
					return
				}

				h.error <- err
				return
			}
		}
	}
}

// Messages to receive message from the receiver. The Message chan may be closed if receiver is no longer available.
func (h *Handler[In, Out]) Messages() <-chan *Message[Out] {
	return h.outgoing
}

// Errors to receive error from the handler.
func (h *Handler[In, Out]) Errors() <-chan error {
	return h.error
}

// handleRecv handles the incoming from the receiver.
// Blocks until receiver.Recv returns. The result from the Recv is either going to Error or incoming:
func (h *Handler[In, Out]) handleRecv(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		m, err := h.receiver.Recv()
		if err != nil {
			select {
			case <-ctx.Done():
				return
			case h.error <- err:
			}
			return
		}
		if m == nil {
			// close the outgoing to signal no more message to be sent, caller should check for whether the
			// chan is closed or not.
			close(h.outgoing)
			return
		}
		select {
		case <-ctx.Done():
			return
		case h.outgoing <- &Message[Out]{
			Msg: m,
		}:
		}
	}
}
