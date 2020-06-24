package resolvers

import (
	"GG-IceCreamShop/api_gateway/clients"
	"GG-IceCreamShop/api_gateway/types"
	"context"
	"proto/auth"
	"proto/ice_cream"

	graphql "github.com/graph-gophers/graphql-go"
)

type RootResolver struct {
	Query
	Mutation
}

type Query struct{}

type CredentialsArgs struct {
	Input types.Credentials
}

func (r *Query) Login(ctx context.Context, args CredentialsArgs) (*AuthResolver, error) {
	payload := &auth.Credentials{
		Email:    args.Input.Email,
		Password: args.Input.Password,
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

func (r *Mutation) CreateIceCream(ctx context.Context, args struct{ Input *types.IceCream }) (*IceCreamResolver, error) {
	payload := &ice_cream.IceCreamDetails{
		Name:                  args.Input.Name,
		ImageClosed:           args.Input.ImageClosed,
		ImageOpen:             args.Input.ImageOpen,
		Description:           args.Input.Description,
		Story:                 args.Input.Story,
		SourcingValues:        *args.Input.SourcingValues,
		Ingredients:           *args.Input.Ingredients,
		AllergyInfo:           *args.Input.AllergyInfo,
		DietaryCertifications: *args.Input.DietaryCertifications,
		ProductId:             string(*args.Input.ProductID),
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

type UpdateIceCreamArgs struct {
	ID    graphql.ID
	Input *types.IceCream
}

func (r *Mutation) UpdateIceCream(ctx context.Context, args UpdateIceCreamArgs) (*IceCreamResolver, error) {
	payload := &ice_cream.IceCreamDetails{
		Id:                    string(args.ID),
		Name:                  args.Input.Name,
		ImageClosed:           args.Input.ImageClosed,
		ImageOpen:             args.Input.ImageOpen,
		Description:           args.Input.Description,
		Story:                 args.Input.Story,
		SourcingValues:        *args.Input.SourcingValues,
		Ingredients:           *args.Input.Ingredients,
		AllergyInfo:           *args.Input.AllergyInfo,
		DietaryCertifications: *args.Input.DietaryCertifications,
		ProductId:             string(*args.Input.ProductID),
	}

	resp, err := clients.IceCream.Update(ctx, payload)

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

func (r *Mutation) DeleteIceCream(ctx context.Context, args struct{ ID graphql.ID }) (*string, error) {
	payload := &ice_cream.IceCreamDetails{
		Id: string(args.ID),
	}

	_, err := clients.IceCream.Delete(ctx, payload)

	return nil, err
}

func (r *Mutation) CreateUser(ctx context.Context, args CredentialsArgs) (*string, error) {
	payload := &auth.Credentials{
		Email:    args.Input.Email,
		Password: args.Input.Password,
	}

	_, err := clients.User.Create(ctx, payload)

	return nil, err
}
