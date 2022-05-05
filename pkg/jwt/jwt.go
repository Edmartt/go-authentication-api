package jwt

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct{
	jwt.StandardClaims
	Attribute string
}

func (cl Claims)GenerateJWT(attribute string, quantity time.Duration) string{
	expirationTime := time.Now().Add(quantity * time.Minute)

	claims := &Claims{
		Attribute: attribute,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := os.Getenv("JWT_SECRET")

	tokenString, tokenError := token.SignedString([]byte(jwtSecret))

	if tokenError != nil{
		log.Println(tokenError)
	}

	return tokenString
}
