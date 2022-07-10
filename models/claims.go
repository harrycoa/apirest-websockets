package models

import "github.com/golang-jwt/jwt"

type AppClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

/*
Go no tiene herencia, pero tiene composicion sobre herencia "jwt.StandardClaims"
*/
