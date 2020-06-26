package main

import (
	"context"
	"proto/ice_cream"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/oklog/ulid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type IceCreamServerRPC struct{}

func (s *IceCreamServerRPC) Get(ctx context.Context, req *ice_cream.IceCreamQuery) (*ice_cream.IceCreams, error) {
	var iceCreams []IceCream
	var totalCount int32
	var hasNext bool
	if err := iceCreamService.GetIceCreams(&iceCreams); err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	if err := iceCreamService.GetIceCreamsCount(&totalCount); err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	if err := iceCreamService.HasNextIceCreams(iceCreams[len(iceCreams)-1].ID, &hasNext); err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	iceCreamsResp := []*ice_cream.IceCreamDetails{}
	for _, iceCream := range iceCreams {
		iceCreamDetails := &ice_cream.IceCreamDetails{
			Id:                    iceCream.ID.String(),
			Name:                  iceCream.Name,
			ImageClosed:           iceCream.ImageClosed,
			ImageOpen:             iceCream.ImageOpen,
			Description:           iceCream.Description,
			Story:                 iceCream.Story,
			SourcingValues:        iceCream.SourcingValues,
			Ingredients:           iceCream.Ingredients,
			AllergyInfo:           iceCream.AllergyInfo,
			DietaryCertifications: iceCream.DietaryCertifications,
			ProductId:             iceCream.ProductID,
		}
		iceCreamsResp = append(iceCreamsResp, iceCreamDetails)
	}

	return &ice_cream.IceCreams{
		IceCreams:  iceCreamsResp,
		TotalCount: totalCount,
		HasNext:    hasNext,
	}, nil
}

func (s *IceCreamServerRPC) Create(ctx context.Context, req *ice_cream.IceCreamDetails) (*ice_cream.IceCreamDetails, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "%v", "unauthenticated access")
	}

	email, ok := md["email"]
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "%v", "unauthenticated access")
	}

	iceCream := &IceCream{
		Name:                  req.Name,
		ImageClosed:           req.ImageClosed,
		ImageOpen:             req.ImageOpen,
		Description:           req.Description,
		Story:                 req.Story,
		SourcingValues:        req.SourcingValues,
		Ingredients:           req.Ingredients,
		AllergyInfo:           req.AllergyInfo,
		DietaryCertifications: req.DietaryCertifications,
		ProductID:             req.ProductId,
		CreatedBy:             email[0],
		UpdatedBy:             email[0],
	}

	if err := iceCream.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	if err := iceCreamService.CreateIceCream(iceCream); err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return &ice_cream.IceCreamDetails{
		Id:                    iceCream.ID.String(),
		Name:                  iceCream.Name,
		ImageClosed:           iceCream.ImageClosed,
		ImageOpen:             iceCream.ImageOpen,
		Description:           iceCream.Description,
		Story:                 iceCream.Story,
		SourcingValues:        iceCream.SourcingValues,
		Ingredients:           iceCream.Ingredients,
		AllergyInfo:           iceCream.AllergyInfo,
		DietaryCertifications: iceCream.DietaryCertifications,
		ProductId:             iceCream.ProductID,
	}, nil
}

func (s *IceCreamServerRPC) Update(ctx context.Context, req *ice_cream.IceCreamDetails) (*ice_cream.IceCreamDetails, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "%v", "unauthenticated access")
	}

	email, ok := md["email"]
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "%v", "unauthenticated access")
	}

	ID, err := ulid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%", err)
	}

	iceCream := &IceCream{
		ID:                    ID,
		Name:                  req.Name,
		ImageClosed:           req.ImageClosed,
		ImageOpen:             req.ImageOpen,
		Description:           req.Description,
		Story:                 req.Story,
		SourcingValues:        req.SourcingValues,
		Ingredients:           req.Ingredients,
		AllergyInfo:           req.AllergyInfo,
		DietaryCertifications: req.DietaryCertifications,
		ProductID:             req.ProductId,
		UpdatedBy:             email[0],
	}

	if err := iceCream.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	if err := iceCreamService.UpdateIceCream(iceCream); err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return &ice_cream.IceCreamDetails{
		Id:                    iceCream.ID.String(),
		Name:                  iceCream.Name,
		ImageClosed:           iceCream.ImageClosed,
		ImageOpen:             iceCream.ImageOpen,
		Description:           iceCream.Description,
		Story:                 iceCream.Story,
		SourcingValues:        iceCream.SourcingValues,
		Ingredients:           iceCream.Ingredients,
		AllergyInfo:           iceCream.AllergyInfo,
		DietaryCertifications: iceCream.DietaryCertifications,
		ProductId:             iceCream.ProductID,
	}, nil
}

func (s *IceCreamServerRPC) Delete(ctx context.Context, req *ice_cream.IceCreamDetails) (*empty.Empty, error) {
	ID, err := ulid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%", err)
	}

	return &empty.Empty{}, iceCreamService.DeleteIceCream(&IceCream{ID: ID})
}
