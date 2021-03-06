package main

import (
	"GG-IceCreamShop/api_gateway/resolvers"

	graphql "github.com/graph-gophers/graphql-go"
)

var RelayHandler = &CustomRelay{
	Schema: graphql.MustParseSchema(GraphSchema, &resolvers.RootResolver{}),
}

var GraphSchema string = `
	type Query {
		login(input: Credentials!): Auth!
		getIceCreams(query: IceCreamQuery): IceCreamResults!
		getIceCreamById(id: ID!): IceCream
	}

	type Mutation {
		createIceCream(input: IceCreamInput!): IceCream
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

	enum SortColumn {
		NAME
		CREATED_AT
		UPDATED_AT
	}

	enum SortDirection {
		ASC
		DESC
	}

	input IceCreamQuery {
		first: Int
		after: ID
		name: String
		sourcing_values: [String!]
		ingredients: [String!]
		sort_column: SortColumn
		sort_direction: SortDirection
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
