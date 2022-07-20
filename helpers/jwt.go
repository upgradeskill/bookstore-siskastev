package helpers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"time"
)

type Claims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

var JWT_SECRET = os.Getenv("JWT_SECRET")

func GenerateJwtToken(email string) (string, error) {
	if JWT_SECRET == "" {
		log.Fatal("[ ERROR ] JWT_SECRET environment variable not provided!\n")
	}

	key := []byte(JWT_SECRET)

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	SignedToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return SignedToken, nil
}

func ValidateToken(strToken string) (*Claims, error) {
	key := []byte(JWT_SECRET)
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(strToken, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return claims, errors.New("invalid token")
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
