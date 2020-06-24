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
		login(input: Credentials!): Auth!
		getIceCreams(query: IceCreamQuery): IceCreamResults!
	}

	type Mutation {
		createIceCream(input: IceCreamInput!): IceCream!
		updateIceCream(id: ID!, input: IceCreamInput!): IceCream!
		deleteIceCream(id: ID!): String
		createUser(input: Credentials!): String
	}

	input Credentials {
		email: String!
		password: String!
	}

	input IceCreamInput {
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

	input IceCreamQuery {
		first: Int
		after: ID
	}

	type Auth {
		jwt_token: String!
	}

	type IceCream {
		id: ID!
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

	type IceCreamResults {
		ice_creams: [IceCream!]
		total_count: Int!
		has_next: Boolean!
	}
`
