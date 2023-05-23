package tokenModel

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	ID       string `json:"id"`
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}
