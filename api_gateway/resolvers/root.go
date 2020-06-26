package resolvers

import (
	"GG-IceCreamShop/api_gateway/clients"
	"GG-IceCreamShop/api_gateway/types"
	"context"
	"fmt"
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
	if err != nil {
		return nil, err
	}

	// hack - graphql-go library doesn't have access to responseWritter to set cookie
	authtoken := ctx.Value("authtoken").(*string)
	*authtoken = resp.JwtToken

	return &AuthResolver{&types.Auth{resp.JwtToken}}, nil
}

func (r *Query) GetIceCreams(ctx context.Context, args struct{ Query *types.IceCreamQuery }) (*IceCreamResultsResolver, error) {
	payload := &ice_cream.IceCreamQuery{}
	if args.Query.First != nil {
		payload.First = int32(*args.Query.First)
	}

	if args.Query.After != nil {
		payload.After = string(*args.Query.After)
	}

	if args.Query.Name != nil {
		payload.Name = string(*args.Query.Name)
	}

	if args.Query.SourcingValues != nil {
		payload.SourcingValues = *args.Query.SourcingValues
	}

	if args.Query.Ingredients != nil {
		payload.Ingredients = *args.Query.Ingredients
	}

	if args.Query.SortColumn != nil {
		value, ok := ice_cream.SortColumn_value[*args.Query.SortColumn]
		if !ok {
			return nil, fmt.Errorf("sort_column value is not a valid enum")
		}
		payload.SortCol = ice_cream.SortColumn(value)
	}

	if args.Query.SortDirection != nil {
		value, ok := ice_cream.SortDir_value[*args.Query.SortDirection]
		if !ok {
			return nil, fmt.Errorf("sort_direction value is not a valid enum")
		}
		payload.SortDir = ice_cream.SortDir(value)
	}

	resp, err := clients.IceCream.Get(ctx, payload)
	if err != nil {
		return nil, err
	}

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

	return &IceCreamResultsResolver{&iceCreamResolvers, resp.TotalCount, resp.HasNext}, err
}

type Mutation struct{}

func (r *Mutation) CreateIceCream(ctx context.Context, args struct{ Input *types.IceCream }) (*IceCreamResolver, error) {
	payload := &ice_cream.IceCreamDetails{
		Name:        args.Input.Name,
		ImageClosed: args.Input.ImageClosed,
		ImageOpen:   args.Input.ImageOpen,
		Description: args.Input.Description,
		Story:       args.Input.Story,
	}

	if args.Input.SourcingValues != nil {
		payload.SourcingValues = *args.Input.SourcingValues
	}

	if args.Input.Ingredients != nil {
		payload.Ingredients = *args.Input.Ingredients
	}

	if args.Input.AllergyInfo != nil {
		payload.AllergyInfo = *args.Input.AllergyInfo
	}

	if args.Input.DietaryCertifications != nil {
		payload.DietaryCertifications = *args.Input.DietaryCertifications
	}

	if args.Input.ProductID != nil {
		payload.ProductId = string(*args.Input.ProductID)
	}

	resp, err := clients.IceCream.Create(ctx, payload)
	if err != nil {
		return nil, err
	}

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
		Id:          string(args.ID),
		Name:        args.Input.Name,
		ImageClosed: args.Input.ImageClosed,
		ImageOpen:   args.Input.ImageOpen,
		Description: args.Input.Description,
		Story:       args.Input.Story,
	}

	if args.Input.SourcingValues != nil {
		payload.SourcingValues = *args.Input.SourcingValues
	}

	if args.Input.Ingredients != nil {
		payload.Ingredients = *args.Input.Ingredients
	}

	if args.Input.AllergyInfo != nil {
		payload.AllergyInfo = *args.Input.AllergyInfo
	}

	if args.Input.DietaryCertifications != nil {
		payload.DietaryCertifications = *args.Input.DietaryCertifications
	}

	if args.Input.ProductID != nil {
		payload.ProductId = string(*args.Input.ProductID)
	}

	resp, err := clients.IceCream.Update(ctx, payload)
	if err != nil {
		return nil, err
	}

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
