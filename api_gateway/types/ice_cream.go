package types

import (
	graphql "github.com/graph-gophers/graphql-go"
)

type IceCreamQuery struct {
	First *float64
	After *graphql.ID
}

type IceCream struct {
	ID                    graphql.ID
	Name                  string
	ImageClosed           string
	ImageOpen             string
	Description           string
	Story                 string
	SourcingValues        *[]string
	Ingredients           *[]string
	AllergyInfo           *string
	DietaryCertifications *string
	ProductID             *graphql.ID
}
