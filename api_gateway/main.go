package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"GG-IceCreamShop/api_gateway/clients"
)

var port string

func init() {
	if port = os.Getenv("PORT"); port == "" {
		port = ":8000"
	}
}

func main() {
	grpcConns := clients.RegisterGrpcServices()
	for _, conn := range grpcConns {
		defer conn.Close()
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/graphiql", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(Graphiql))
	})

	mux.Handle("/query", RelayHandler)

	handlers := SetMiddlewares(mux)

	server := &http.Server{
		Addr:    port,
		Handler: handlers,
	}

	go func() {
		log.Fatalf("listen err: %v", server.ListenAndServe())
	}()

	log.Printf("graphql server listening at %s", server.Addr)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-shutdown

	log.Printf("shutdown command received, shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server.Shutdown(ctx)

	log.Printf("server shutdown completed properly")
}
