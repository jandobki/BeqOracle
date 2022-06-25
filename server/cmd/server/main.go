package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/jandobki/beqoracle/server/internal/model"
	"github.com/jandobki/beqoracle/server/internal/server"
	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterBeqOracleServer(grpcServer, server.NewServer())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		log.Println("stopping server")
		grpcServer.GracefulStop()
	}()

	log.Println("starting server")
	err = grpcServer.Serve(l)
	if err != nil {
		log.Fatalf("cannot serve: %v", err)
	}
	log.Println("server stopped")
}
