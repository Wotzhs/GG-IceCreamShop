package main

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

var (
	chainedUnaryInterceptors grpc.ServerOption
	passwordCensorRe         *regexp.Regexp
)

func init() {
	chainedUnaryInterceptors = grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		RequestLogger,
		RequestTimer,
		ResponseLogger,
	))

	passwordCensorRe = regexp.MustCompile(`(?i)(password:)(\"\w+\")`)
}

func RequestLogger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logInfo := fmt.Sprintf("%v %v\n", info.FullMethod, req)
	censoredLogInfo := passwordCensorRe.ReplaceAllString(logInfo, `$1"****"`)
	log.Printf(censoredLogInfo)
	h, err := handler(ctx, req)
	return h, err
}

func RequestTimer(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	h, err := handler(ctx, req)
	logInfo := fmt.Sprintf("%v %v completed in %.2fms", info.FullMethod, req, float64(time.Now().Sub(start).Nanoseconds())/float64(time.Millisecond))
	censoredLogInfo := passwordCensorRe.ReplaceAllString(logInfo, `$1"****"`)
	log.Printf(censoredLogInfo)
	return h, err
}

func ResponseLogger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	h, err := handler(ctx, req)
	logInfo := fmt.Sprintf("%v %v response: %v err: %v", info.FullMethod, req, h, err)
	censoredLogInfo := passwordCensorRe.ReplaceAllString(logInfo, `$1"****"`)
	log.Printf(censoredLogInfo)
	return h, err
}
