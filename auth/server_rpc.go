package main

import (
	"context"
	"proto/auth"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServerRPC struct{}

func (s *AuthServerRPC) GenerateJWTToken(ctx context.Context, req *auth.Credentials) (*auth.JWTToken, error) {
	resp, err := UserClient.GetPasswordHash(ctx, &auth.Credentials{Email: req.Email})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%v", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(resp.Hash), []byte(req.Password)); err != nil {
		return nil, status.Errorf(codes.NotFound, "%v", err)
	}

	token, err := authService.GenerateToken(req.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return &auth.JWTToken{JwtToken: token}, nil
}
