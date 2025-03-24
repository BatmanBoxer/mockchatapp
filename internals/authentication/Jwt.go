package auth

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("batmanboxer")

func ValidateJwt(tokenString string) (*jwt.StandardClaims, error) {
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid token format")
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return secretKey, nil
		})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if ok && token.Valid && claims.Issuer == "batmanissuer" {
		return claims, nil
	}

	return nil, errors.New("invalid token signature")
}

func GenerateJwt(id string) (string, error) {

	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		Issuer:    "batmanissuer",
		Id:        id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
  log.Println(id)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
