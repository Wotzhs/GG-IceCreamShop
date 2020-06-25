package main

import (
	"log"
	"os"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
)

var (
	authService  *AuthService
	jwtSecretKey string
)

func init() {
	if jwtSecretKey = os.Getenv("JWT_SECRET_KEY"); jwtSecretKey == "" {
		log.Fatalf("env JWT_SECRET_KEY not set")
	}
}

type AuthService struct{}

func (s *AuthService) GenerateToken(email string) (string, error) {
	claims := jws.Claims{}
	claims.Set("email", email)
	claims.SetExpiration(time.Now().Add(24 * time.Hour))
	claims.SetIssuedAt(time.Now())

	jwt := jws.NewJWT(claims, crypto.SigningMethodHS256)

	token, err := jwt.Serialize([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}

	return string(token), nil
}
