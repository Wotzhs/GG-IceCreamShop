module GG-IceCreamShop

go 1.14

require (
	github.com/golang/protobuf v1.4.2
	github.com/graph-gophers/graphql-go v0.0.0-20200622220639-c1d9693c95a6
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0 // indirect
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.24.0
	proto/auth v0.0.0-00010101000000-000000000000
	proto/ice_cream v0.0.0-00010101000000-000000000000
	proto/user v0.0.0-00010101000000-000000000000
)

replace proto/auth => ./proto/auth

replace proto/ice_cream => ./proto/ice_cream

replace proto/user => ./proto/user
