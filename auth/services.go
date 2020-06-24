package main

import (
	"GG-IceCreamShop/proto/auth"
	"context"
)

type AuthService struct{}

func (s *AuthService) GenerateJWTToken(ctx context.Context, req *auth.Credentials) (*auth.JWTToken, error) {
	return &auth.JWTToken{
		JwtToken: "hello word",
	}, nil
}
