package types

type Credentials struct {
	Email    string
	Password string
}

type Auth struct {
	JWTToken string
}
