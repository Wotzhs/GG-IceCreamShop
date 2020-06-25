package main

import (
	"log"
	"os"
	"proto/user"

	"google.golang.org/grpc"
)

var (
	UserClient     user.UserClient
	userClientAddr string
)

func init() {
	if userClientAddr = os.Getenv("USER_CLIENT_ADDR"); userClientAddr == "" {
		userClientAddr = "0.0.0.0:50053"
	}
}

func RegisterGrpcServices() []*grpc.ClientConn {
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	UserConn, err := grpc.Dial(userClientAddr, dialOpts...)
	if err != nil {
		log.Fatalf("failed to connect to user service err: %v", err)
	}

	UserClient = user.NewUserClient(UserConn)

	return []*grpc.ClientConn{UserConn}
}
