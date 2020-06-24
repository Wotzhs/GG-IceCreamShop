package resolvers

import (
	"GG-IceCreamShop/api_gateway/types"
)

type AuthResolver struct {
	a *types.Auth
}

func (r *AuthResolver) JwtToken() string {
	return r.a.JWTToken
}
