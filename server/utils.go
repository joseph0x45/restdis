package main

import (
	"github.com/golang-jwt/jwt/v5"
)

func GenJWT(username string, jwt_secret string) (token string, err error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})
	token, err = t.SignedString([]byte(jwt_secret))
  return
}
