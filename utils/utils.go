package utils

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
    UserID uint `json:"user_id"` 
    jwt.StandardClaims
}

var jwtKey = []byte("supersecretkey")

func GenerateToken(userID string) (string, error) {
    
    id, err := strconv.ParseUint(userID, 10, 32) 
    if err != nil {
        return "", errors.New("userID tidak valid")
    }

    expirationTime := time.Now().Add(24 * time.Hour) 
    claims := &Claims{
        UserID: uint(id), 
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

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