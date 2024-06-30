package utils

import (
    "time"

    "github.com/dgrijalva/jwt-go"
)

// GenerateJWT generates a JWT token with user details as claims
func GenerateJWT(userID int, name, email string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "userID": userID,
        "name":   name,
        "email":  email,
        "exp":    time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, err := token.SignedString([]byte("your_secret_key"))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

// VerifyJWT verifies a JWT token
func VerifyJWT(tokenString string) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte("your_secret_key"), nil
    })

    if err != nil {
        return nil, err
    }

    return token, nil
}
