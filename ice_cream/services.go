package main

import (
	"GG-IceCreamShop/proto/ice_cream"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
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
