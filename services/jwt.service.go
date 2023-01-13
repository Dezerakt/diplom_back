package services

import (
	"diplom_back/initializers"
	"github.com/golang-jwt/jwt"
	"log"
)

type Jwt struct {
	config         initializers.Config
	secret         string
	expirationTime int64
}

func NewJWT() Jwt {
	config, _ := initializers.LoadConfig(".")

	return Jwt{
		config:         config,
		secret:         config.JwtSecretKey,
		expirationTime: config.TokenTTL * 3600,
	}
}

func (object *Jwt) GetNewAccessToken() string {
	tokenClaimsObject := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.StandardClaims{
		ExpiresAt: object.expirationTime,
		Id:        "3",
		Issuer:    "localhost",
	})
	tokenString, err := tokenClaimsObject.SignedString([]byte(object.secret))
	if err != nil {
		log.Fatal("error while generate access token", err)
	}
	return tokenString
}

/*
func (object *Jwt) getNewRefreshToken() string {

}*/
