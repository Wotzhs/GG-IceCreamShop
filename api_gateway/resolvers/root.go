package resolvers

import (
	"GG-IceCreamShop/api_gateway/clients"
	"GG-IceCreamShop/api_gateway/types"
	"GG-IceCreamShop/proto/auth"
	"GG-IceCreamShop/proto/ice_cream"
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

type RootResolver struct{}

func (r *RootResolver) Login(ctx context.Context, args struct{ Email, Password string }) (*AuthResolver, error) {
	payload := &auth.Credentials{
		Email:    args.Email,
		Password: args.Password,
	}

	resp, err := clients.Auth.GenerateJWTToken(ctx, payload)
	return &AuthResolver{&types.Auth{resp.JwtToken}}, err
}

func (r *RootResolver) GetIceCreams(ctx context.Context, args struct{ Query *types.IceCreamQuery }) (*IceCreamResultsResolver, error) {
	payload := &ice_cream.IceCreamQuery{
		First: int32(*args.Query.First),
		After: string(*args.Query.After),
	}

	resp, err := clients.IceCream.Get(ctx, payload)

	iceCreamResolvers := []*IceCreamResolver{}

	for _, iceCream := range resp.IceCreams {
		iceCreamType := types.IceCream{
			ID:                    graphql.ID(iceCream.Id),
			Name:                  iceCream.Name,
			ImageClosed:           iceCream.ImageClosed,
			ImageOpen:             iceCream.ImageOpen,
			Description:           iceCream.Description,
			Story:                 iceCream.Story,
			SourcingValues:        &iceCream.SourcingValues,
			Ingredients:           &iceCream.Ingredients,
			AllergyInfo:           iceCream.AllergyInfo,
			DietaryCertifications: iceCream.DietaryCertifications,
			ProductID:             graphql.ID(iceCream.ProductId),
		}
		iceCreamResolvers = append(iceCreamResolvers, &IceCreamResolver{&iceCreamType})
	}

	return &IceCreamResultsResolver{&iceCreamResolvers, 10, false}, err
}
