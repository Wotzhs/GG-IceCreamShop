package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"proto/ice_cream"

	"google.golang.org/grpc"
)

var port string

func init() {
	if port = os.Getenv("PORT"); port == "" {
		port = ":50052"
	}
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on port %s", port)
	}

	opts := []grpc.ServerOption{chainedUnaryInterceptors}
	server := grpc.NewServer(opts...)
	iceCreamServerRPC := &IceCreamServerRPC{}
	ice_cream.RegisterIceCreamServer(server, iceCreamServerRPC)

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatalf("failed to start ice cream service err: %v", err)
		}
	}()

	log.Printf("ice cream service listening at %s", port)

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
