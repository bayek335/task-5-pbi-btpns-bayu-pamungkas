package app

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}
