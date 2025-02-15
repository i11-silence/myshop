package service

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v5"
	"myshop/model"
	"strings"
	"time"
)

var TOKENKEY = []byte("asdfghj")

const TOKEN_EXP = 1 * time.Hour
const REFRESH_TOKEN_EXP = 72 * time.Hour

func GenerateToken(username string) (string, error) {
	claim := model.Token{
		username,
		false,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TOKEN_EXP)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	tokenobj := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token1, err := tokenobj.SignedString(TOKENKEY)
	return token1, err
}

func GenerateRefresh_Token(username string) (string, error) {
	claim := model.Token{
		username,
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(REFRESH_TOKEN_EXP)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	refresh_tokenobj := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	refresh_token, err := refresh_tokenobj.SignedString(TOKENKEY)
	return refresh_token, err
}

func ValidateToken(c *app.RequestContext) (string, error) {
	token1 := string(c.GetHeader("Authorization"))
	if token1 == "" {
		return "", nil
	}
	parts := strings.SplitN(token1, " ", 2)
	token1 = parts[1]
	tokenobj, err := jwt.ParseWithClaims(token1, &model.Token{}, func(_ *jwt.Token) (interface{}, error) {
		return TOKENKEY, nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := tokenobj.Claims.(*model.Token)
	if ok && tokenobj.Valid {
		return claims.Username, nil
	} else {
		return "", err
	}
}

func ValidaterefreshToken(c *app.RequestContext) (string, error) {
	refresh_token := c.Query("refresh_token")
	if refresh_token == "" {
		return "", nil
	}
	refresh_tokenobj, err := jwt.ParseWithClaims(refresh_token, &model.Token{}, func(_ *jwt.Token) (interface{}, error) {
		return TOKENKEY, nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := refresh_tokenobj.Claims.(*model.Token)
	if ok && refresh_tokenobj.Valid && claims.IsRefresh {
		return claims.Username, nil
	} else {
		return "", err
	}
}
