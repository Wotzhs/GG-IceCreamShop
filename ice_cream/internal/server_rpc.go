package internal

import (
	"context"
	"proto/ice_cream"

	"GG-IceCreamShop/ice_cream/internal/models"
	"GG-IceCreamShop/ice_cream/internal/services"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/oklog/ulid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type IceCreamServerRPC struct{}

func (s *IceCreamServerRPC) Get(ctx context.Context, req *ice_cream.IceCreamQuery) (*ice_cream.IceCreams, error) {
	var iceCreams []models.IceCream
	var totalCount int32
	var hasNext bool
	if err := services.IceCream.GetIceCreams(req, &iceCreams); err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	if err := services.IceCream.GetIceCreamsCount(req, &totalCount); err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	if len(iceCreams) > 0 {
		if err := services.IceCream.HasNextIceCreams(iceCreams[len(iceCreams)-1].ID, &hasNext); err != nil {
			return nil, status.Errorf(codes.Internal, "%v", err)
		}
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

func (s *IceCreamServerRPC) GetById(ctx context.Context, req *ice_cream.IceCreamQuery) (*ice_cream.IceCreamDetails, error) {
	var iceCream models.IceCream
	ID, err := ulid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	if err := services.IceCream.GetIceCreamByID(ID, &iceCream); err != nil {
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

func (s *IceCreamServerRPC) Create(ctx context.Context, req *ice_cream.IceCreamDetails) (*ice_cream.IceCreamDetails, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "%v", "unauthenticated access")
	}

	email, ok := md["email"]
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "%v", "unauthenticated access")
	}

	iceCream := &models.IceCream{
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

	if err := services.IceCream.CreateIceCream(iceCream); err != nil {
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

	iceCream := &models.IceCream{
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

	if err := services.IceCream.UpdateIceCream(iceCream); err != nil {
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

	return &empty.Empty{}, services.IceCream.DeleteIceCream(&models.IceCream{ID: ID})
}
