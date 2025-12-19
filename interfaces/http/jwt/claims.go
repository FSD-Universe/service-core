// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

//go:build http && httpjwt

package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"half-nothing.cn/service-core/interfaces/config"
	"half-nothing.cn/service-core/interfaces/database/entity"
)

type Claims struct {
	Uid        uint   `json:"uid"`
	Cid        uint   `json:"cid"`
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

type ClaimFactoryInterface interface {
	EmptyClaim() *Claims
	EmptyFsdClaim() *FsdClaims
	CreateClaim(user *entity.User, flushToken bool) *Claims
	CreateFsdClaim(user *entity.User) *FsdClaims
	GenerateKey(claims jwt.Claims) (string, error)
	VerifyJwt(jwtString string, claim jwt.Claims) (bool, error)
	GetJWTConfig() *config.JwtConfig
}

type ContentSetter interface {
	SetUid(uid uint)
	SetCid(cid uint)
	SetPermission(permission uint64)
	SetRating(rating int)
	SetRaw(raw *Claims)
}

type Content struct {
	Uid        uint
	Cid        uint
	Permission uint64
	Rating     int
	Raw        *Claims
}

func (c *Content) SetUid(uid uint) {
	c.Uid = uid
}

func (c *Content) SetCid(cid uint) {
	c.Cid = cid
}

func (c *Content) SetPermission(permission uint64) {
	c.Permission = permission
}

func (c *Content) SetRating(rating int) {
	c.Rating = rating
}

func (c *Content) SetRaw(raw *Claims) {
	c.Raw = raw
}

func SetJwtContent[T ContentSetter](data T, ctx echo.Context) error {
	claim, ok := ctx.Get("user").(*Claims)
	if !ok {
		return errors.New("JWT token not found in context")
	}
	data.SetPermission(claim.Permission)
	data.SetUid(claim.Uid)
	data.SetCid(claim.Cid)
	data.SetRating(claim.Rating)
	data.SetRaw(claim)
	return nil
}
