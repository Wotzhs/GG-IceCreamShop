package resolvers

import (
	"GG-IceCreamShop/api_gateway/types"

	graphql "github.com/graph-gophers/graphql-go"
)

type IceCreamResolver struct {
	i *types.IceCream
}

func (r *IceCreamResolver) ID() graphql.ID {
	return r.i.ID
}

func (r *IceCreamResolver) Name() string {
	return r.i.Name
}

func (r *IceCreamResolver) ImageClosed() string {
	return r.i.ImageClosed
}

func (r *IceCreamResolver) ImageOpen() string {
	return r.i.ImageOpen
}

func (r *IceCreamResolver) Description() string {
	return r.i.Description
}

func (r *IceCreamResolver) Story() string {
	return r.i.Story
}

func (r *IceCreamResolver) SourcingValues() *[]string {
	return r.i.SourcingValues
}

func (r *IceCreamResolver) Ingredients() *[]string {
	return r.i.Ingredients
}

func (r *IceCreamResolver) AllergyInfo() *string {
	return r.i.AllergyInfo
}

func (r *IceCreamResolver) DietaryCertifications() *string {
	return r.i.DietaryCertifications
}

func (r *IceCreamResolver) ProductID() *graphql.ID {
	return r.i.ProductID
}

type IceCreamResultsResolver struct {
	iceCreams  *[]*IceCreamResolver
	totalCount int32
	hasNext    bool
}

func (r *IceCreamResultsResolver) IceCreams() *[]*IceCreamResolver {
	return r.iceCreams
}

func (r *IceCreamResultsResolver) TotalCount() int32 {
	return r.totalCount
}

func (r *IceCreamResultsResolver) HasNext() bool {
	return r.hasNext
}
