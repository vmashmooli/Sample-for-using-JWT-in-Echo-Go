package main

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// Structure for token
type (
	Token struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		jwt.StandardClaims
	}
)

// A function to create token by User Id and User Name with configuration
func CreateToken(id, name string) (string, error) {
	claims := Token{
		ID:             id,
		Name:           name,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Duration(config.Setting.JWT.Age) * time.Hour).Unix()},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Setting.JWT.Secret))
	if err != nil {
		return "", errors.New("Error in creating Token")
	}
	return t, nil
}

// A function to check token validity
func CheckToken(token string) (string, string, error) {
	claims := jwt.MapClaims{}
	t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Setting.JWT.Secret), nil
	})
	if !t.Valid || err != nil {
		return "", "", err
	}
	return claims["id"].(string), claims["name"].(string), nil
}
