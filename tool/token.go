package tool

import (
	"aurora/blog/api/entity"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

func GenerateToken(userName string) (string, error) {
	expire := viper.GetInt("authorization.token.expire")
	secret := viper.GetString("authorization.token.secret")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, entity.CustomClaim{
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expire) * time.Second).Unix(),
		},
	})

	return token.SignedString([]byte(secret))
}

func ParseToken(token string) (*entity.CustomClaim, error) {
	secret := viper.GetString("authorization.token.secret")

	authToken, err := jwt.ParseWithClaims(token, &entity.CustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claim, ok := authToken.Claims.(*entity.CustomClaim); ok && authToken.Valid {
		return claim, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
