package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/jandobki/beqoracle/server/internal/event"
	pb "github.com/jandobki/beqoracle/server/internal/model"
	"github.com/jandobki/beqoracle/server/internal/server"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	event.InitDB(ctx, "beqoracle")

	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterBeqOracleServer(grpcServer, server.NewServer(ctx))

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		log.Println("stopping...")
		cancel()
		grpcServer.GracefulStop()
	}()

	log.Println("serving...")

	err = grpcServer.Serve(l)

	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("stopped")
}
