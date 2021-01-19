package entity

import "github.com/dgrijalva/jwt-go"

type CustomClaim struct {
	UserName string
	jwt.StandardClaims
}
