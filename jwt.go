package main

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

var jwt_key = []byte("my_secret_key")

func CreateJwtToken(sub string, dur time.Duration) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "cdbb.me",
		Subject:   sub,
		Audience:  nil,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(dur)),
		NotBefore: nil,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        uuid.New().String(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(jwt_key)
	if err != nil {
		return ""
	} else {
		return tokenString
	}
}

var keyFunc = func(token *jwt.Token) (interface{}, error) {
	// Don't forget to validate the alg is what you expect:
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	return jwt_key, nil
}

func ParseToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, keyFunc)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
