package main

import (
	"context"
	"proto/ice_cream"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type IceCreamService struct{}

func (s *IceCreamService) Get(ctx context.Context, req *ice_cream.IceCreamQuery) (*ice_cream.IceCreams, error) {
	return &ice_cream.IceCreams{
		IceCreams: []*ice_cream.IceCreamDetails{
			&ice_cream.IceCreamDetails{
				Id:                    "Hello World",
				Name:                  "Hello World",
				ImageClosed:           "Hello World",
				ImageOpen:             "Hello World",
				Description:           "Hello World",
				Story:                 "Hello World",
				SourcingValues:        []string{"Hello World"},
				Ingredients:           []string{"Hello World"},
				AllergyInfo:           "Hello World",
				DietaryCertifications: "Hello World",
				ProductId:             "Hello World",
			},
		},
	}, nil
}

func (s *IceCreamService) Create(ctx context.Context, req *ice_cream.IceCreamDetails) (*ice_cream.IceCreamDetails, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "%v", "unauthenticated access")
	}

	_, ok = md["email"]
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "%v", "unauthenticated access")
	}

	return &ice_cream.IceCreamDetails{
		Id: "Hello World",
	}, nil
}

func (s *IceCreamService) Update(ctx context.Context, req *ice_cream.IceCreamDetails) (*ice_cream.IceCreamDetails, error) {
	return &ice_cream.IceCreamDetails{
		Id: "Hello World",
	}, nil
}

func (s *IceCreamService) Delete(ctx context.Context, req *ice_cream.IceCreamDetails) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
