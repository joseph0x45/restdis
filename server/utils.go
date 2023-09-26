package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func GenJWT(username string, jwt_secret string) (token string, err error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})
	token, err = t.SignedString([]byte(jwt_secret))
	return
}

func UserAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth_token := r.Header.Get("Authorization")
		if auth_token == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		token, err := jwt.Parse(auth_token, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("")
			}
			var config = Config{}
			err := Db.Get(&config, `select from config where id=main`)
			if err != nil {
				return nil, fmt.Errorf("")
			}
			return []byte(config.JwtSecret), nil
		})
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			rights := claims["rights"].(string)
      user := claims["user"].(string)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
	})
}

func Auth(constraint string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth_token := r.Header.Get("Authorization")
			if auth_token == "" {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			token, err := jwt.Parse(auth_token, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("")
				}
				var config = Config{}
				err := Db.Get(&config, `select from config where id=main`)
				if err != nil {
					return nil, fmt.Errorf("")
				}
				return []byte(config.JwtSecret), nil
			})
			if err != nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				rights := claims["rights"].(string)
				if constraint == "write" {
					if !strings.Contains(rights, "w") {
						http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
						return
					}
				}
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		})

	}
}
