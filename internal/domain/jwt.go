package domain

import "github.com/golang-jwt/jwt/v5"

// Struct untuk payload JWT (Claims)
type Claims struct {
    Username string `json:"username"`
    jwt.RegisteredClaims
}