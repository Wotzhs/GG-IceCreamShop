package main

import (
	"context"
	"proto/auth"

	"github.com/golang/protobuf/ptypes/empty"
)

type UserService struct{}

func (s *UserService) Create(ctx context.Context, req *auth.Credentials) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
