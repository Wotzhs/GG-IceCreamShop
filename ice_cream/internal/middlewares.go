package internal

import (
	"context"
	"log"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

var chainedUnaryInterceptors grpc.ServerOption

func init() {
	chainedUnaryInterceptors = grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		RequestLogger,
		RequestTimer,
		ResponseLogger,
	))
}

func RequestLogger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("%v %v\n", info.FullMethod, req)
	h, err := handler(ctx, req)
	return h, err
}

func RequestTimer(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	h, err := handler(ctx, req)
	log.Printf("%v %v completed in %.2fms", info.FullMethod, req, float64(time.Now().Sub(start).Nanoseconds())/float64(time.Millisecond))
	return h, err
}

func ResponseLogger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	h, err := handler(ctx, req)
	log.Printf("%v %v response: %v err: %v", info.FullMethod, req, h, err)
	return h, err
}
