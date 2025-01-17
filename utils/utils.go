package utils

import (
    "errors"
    "github.com/golang-jwt/jwt/v4"
)

type Claims struct {
    UserID uint `json:"user_id"`
    jwt.StandardClaims
}

var jwtKey = []byte("your_secret_key")

func VerifyToken(tokenString string) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil || !token.Valid {
        return nil, errors.New("token tidak valid")
    }

    return claims, nil
}
