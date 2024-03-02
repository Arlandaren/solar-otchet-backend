package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	// Role   string `json:"role"`   	решено убрать до лучших времен
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}
