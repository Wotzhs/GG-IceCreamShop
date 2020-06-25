package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"

	"proto/user"
)

var port string

func init() {
	if port = os.Getenv("PORT"); port == "" {
		port = ":50053"
	}
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on port %s", port)
	}

	opts := []grpc.ServerOption{chainedUnaryInterceptors}
	server := grpc.NewServer(opts...)
	userServerRPC := &UserServerRPC{}
	user.RegisterUserServer(server, userServerRPC)

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatalf("failed to start auth service err: %v", err)
		}
	}()

	log.Printf("auth service listening at %s", port)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	log.Printf("shutdown command received, shutting down server")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.GracefulStop()
	listener.Close()

	log.Printf("server shutdown completed properly")
}
