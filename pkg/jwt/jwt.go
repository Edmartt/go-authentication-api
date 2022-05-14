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
		return
	}

	return 
}

func (wrap *JWTWrapper) ValidateToken(signedToken string) (claims *Claims, jwtError error){
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(wrap.SecretKey), nil
		},
	)

	if err != nil{
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok{
		err = errors.New("couln't parse with claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix(){
		err = errors.New("JWT expired")
		return
	}

	return tokenString
}
