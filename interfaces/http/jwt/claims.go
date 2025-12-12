//go:build http && httpjwt

package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
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

type ContentSetter interface {
	SetUid(uid uint)
	SetCid(cid int)
	SetPermission(permission uint64)
	SetRating(rating int)
}

func SetJwtContent[T ContentSetter](data T, ctx echo.Context) error {
	token, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return errors.New("JWT token not found in context")
	}
	claim, ok := token.Claims.(*Claims)
	if !ok {
		return errors.New("invalid claim type")
	}
	data.SetPermission(claim.Permission)
	data.SetUid(claim.Uid)
	data.SetCid(claim.Cid)
	data.SetRating(claim.Rating)
	return nil
}
