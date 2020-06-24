package main

import (
	"GG-IceCreamShop/api_gateway/resolvers"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

var RelayHandler = &relay.Handler{
	Schema: graphql.MustParseSchema(GraphSchema, &resolvers.RootResolver{}),
}

var GraphSchema string = `
	type Query {
		login(email: String!, password: String!): Auth!
		getIceCreams(query: IceCreamQuery): IceCreamResults!
	}

	input IceCreamQuery {
		first: Int
		after: ID
	}

	type Auth {
		jwt_token: String!
	}

	type IceCreamResults {
		ice_creams: [IceCream!]
		total_count: Int!
		has_next: Boolean!
	}

	type Mutation {
		createIceCream(details: NewIceCream!): IceCream!
	#	updateIceCream(details: UpdateIceCream): IceCream!
	#	deleteIceCream(id: ID): String
	}

	input NewIceCream {
		name: String!
		image_closed: String!
		image_open: String!
		description: String!
		story: String!
		sourcing_values: [String!]
		ingredients: [String!]
		allergy_info: String
		dietary_certifications: String
		product_id: ID
	}

	input UpdateIceCream {
		id: ID!
		updated_details: NewIceCream
	}

	type IceCream {
		id: ID
		name: String!
		image_closed: String!
		image_open: String!
		description: String!
		story: String!
		sourcing_values: [String!]
		ingredients: [String!]
		allergy_info: String
		dietary_certifications: String
		product_id: ID
	}
`
