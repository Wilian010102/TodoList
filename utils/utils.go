package utils

import (
    "errors"
    "time"
    "github.com/golang-jwt/jwt/v4"
    "log"
)

// Secret key for JWT
var jwtKey = []byte("supersecretkey")

type Claims struct {
    UserID uint `json:"user_id"`
    jwt.RegisteredClaims
}

func VerifyToken(tokenString string) (*Claims, error) {
    claims := &Claims{}

    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        log.Println("JWT Parsing Error:", err)
        return nil, errors.New("token not valid")
    }

    if !token.Valid {
        log.Println("JWT Token not valid")
        return nil, errors.New("token not valid")
    }

    if time.Now().Unix() > claims.ExpiresAt.Time.Unix() {
        log.Println("JWT Token expired")
        return nil, errors.New("token expired")
    }

    log.Println("JWT Valid for user_id:", claims.UserID)
    return claims, nil
}
