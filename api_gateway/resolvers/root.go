package resolvers

import (
	"GG-IceCreamShop/api_gateway/clients"
	"GG-IceCreamShop/api_gateway/types"
	"GG-IceCreamShop/proto/auth"
	"GG-IceCreamShop/proto/ice_cream"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

type RootResolver struct {
	Query
	Mutation
}

type Query struct{}

func (r *Query) Login(ctx context.Context, args struct{ Email, Password string }) (*AuthResolver, error) {
	payload := &auth.Credentials{
		Email:    args.Email,
		Password: args.Password,
	}

	resp, err := clients.Auth.GenerateJWTToken(ctx, payload)
	return &AuthResolver{&types.Auth{resp.JwtToken}}, err
}

func (r *Query) GetIceCreams(ctx context.Context, args struct{ Query *types.IceCreamQuery }) (*IceCreamResultsResolver, error) {
	payload := &ice_cream.IceCreamQuery{
		First: int32(*args.Query.First),
		After: string(*args.Query.After),
	}

	resp, err := clients.IceCream.Get(ctx, payload)

	iceCreamResolvers := []*IceCreamResolver{}

	for _, iceCream := range resp.IceCreams {
		productId := graphql.ID(iceCream.ProductId)
		iceCreamType := types.IceCream{
			ID:                    graphql.ID(iceCream.Id),
			Name:                  iceCream.Name,
			ImageClosed:           iceCream.ImageClosed,
			ImageOpen:             iceCream.ImageOpen,
			Description:           iceCream.Description,
			Story:                 iceCream.Story,
			SourcingValues:        &iceCream.SourcingValues,
			Ingredients:           &iceCream.Ingredients,
			AllergyInfo:           &iceCream.AllergyInfo,
			DietaryCertifications: &iceCream.DietaryCertifications,
			ProductID:             &productId,
		}
		iceCreamResolvers = append(iceCreamResolvers, &IceCreamResolver{&iceCreamType})
	}

	return &IceCreamResultsResolver{&iceCreamResolvers, 10, false}, err
}

type Mutation struct{}

func (r *Mutation) CreateIceCream(ctx context.Context, args struct{ Details *types.IceCream }) (*IceCreamResolver, error) {
	payload := &ice_cream.IceCreamDetails{
		Name:                  args.Details.Name,
		ImageClosed:           args.Details.ImageClosed,
		ImageOpen:             args.Details.ImageOpen,
		Description:           args.Details.Description,
		Story:                 args.Details.Story,
		SourcingValues:        *args.Details.SourcingValues,
		Ingredients:           *args.Details.Ingredients,
		AllergyInfo:           *args.Details.AllergyInfo,
		DietaryCertifications: *args.Details.DietaryCertifications,
		ProductId:             string(*args.Details.ProductID),
	}

	resp, err := clients.IceCream.Create(ctx, payload)

	productID := graphql.ID(resp.ProductId)

	iceCream := &types.IceCream{
		ID:                    graphql.ID(resp.Id),
		Name:                  resp.Name,
		ImageClosed:           resp.ImageClosed,
		ImageOpen:             resp.ImageOpen,
		Description:           resp.Description,
		Story:                 resp.Story,
		SourcingValues:        &resp.SourcingValues,
		Ingredients:           &resp.Ingredients,
		AllergyInfo:           &resp.AllergyInfo,
		DietaryCertifications: &resp.DietaryCertifications,
		ProductID:             &productID,
	}

	return &IceCreamResolver{iceCream}, err
}
