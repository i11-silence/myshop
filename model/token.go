package model

import "github.com/golang-jwt/jwt/v5"

type Token struct {
	Username  string `json:"username"`
	IsRefresh bool   `json:"is_refresh"`
	jwt.RegisteredClaims
}
