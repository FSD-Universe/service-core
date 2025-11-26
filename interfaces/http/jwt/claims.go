//go:build httpjwt

package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Uid        uint   `json:"uid"`
	Cid        int    `json:"cid"`
	Username   string `json:"username"`
	Permission uint64 `json:"permission"`
	Rating     int    `json:"rating"`
	FlushToken bool   `json:"flushToken"`
	jwt.RegisteredClaims
}

type FsdClaims struct {
	ControllerRating int `json:"controller_rating"`
	PilotRating      int `json:"pilot_rating"`
	jwt.RegisteredClaims
}

type ClaimsFactoryInterface interface {
	CreateClaims(uid uint, cid int, username string, permission uint64, rating int, flushToken bool) *Claims
	CreateFsdClaims(controllerRating int, pilotRating int) *FsdClaims
	GenerateKey(claims *jwt.RegisteredClaims) (string, error)
}
