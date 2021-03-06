package main

import (
	"context"
	"proto/auth"
	"proto/user"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServerRPC struct{}

func (s *UserServerRPC) Create(ctx context.Context, req *auth.Credentials) (*empty.Empty, error) {
	user := &User{
		Email:    req.Email,
		Password: req.Password,
	}

	if err := user.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	if err := user.HashPassword(); err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	if err := userService.CreateUser(user); err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return &empty.Empty{}, nil
}

func (s *UserServerRPC) GetPasswordHash(ctx context.Context, req *auth.Credentials) (*user.PasswordHash, error) {
	hash, err := userService.GetPasswordHash(req.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}
	return &user.PasswordHash{Hash: hash}, nil
}
