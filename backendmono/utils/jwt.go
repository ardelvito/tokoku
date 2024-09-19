package utils

import (
    "github.com/golang-jwt/jwt/v4"
    "time"
)

func GenerateJWT(userID int, secret string) (string, error) {
    claims := jwt.MapClaims{
        "id":  userID,
        "exp": time.Now().Add(time.Hour * 72).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}
