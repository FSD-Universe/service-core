// Package jwt
package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	c "half-nothing.cn/service-core/interfaces/config"
	"half-nothing.cn/service-core/interfaces/database/entity"
	. "half-nothing.cn/service-core/interfaces/http/jwt"
)

type ClaimFactory struct {
	config *c.JwtConfig
}

func NewClaimFactory(config *c.JwtConfig) *ClaimFactory {
	return &ClaimFactory{
		config: config,
	}
}

func (f *ClaimFactory) CreateClaims(user *entity.User, flushToken bool) *Claims {
	var expiredDuration time.Duration
	if flushToken {
		expiredDuration = f.config.RefreshExpireDuration
	} else {
		expiredDuration = f.config.ExpireDuration
	}
	now := time.Now()
	return &Claims{
		Uid:        user.ID,
		Cid:        user.Cid,
		Username:   user.Username,
		Permission: user.Permission,
		Rating:     user.Rating,
		FlushToken: flushToken,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    f.config.Issuer,
			Subject:   user.Username,
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(expiredDuration)),
		},
	}
}

func (f *ClaimFactory) CreateFsdClaims(user *entity.User) *FsdClaims {
	now := time.Now()
	return &FsdClaims{
		ControllerRating: user.Rating,
		PilotRating:      user.Rating,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    f.config.Issuer,
			Subject:   user.Username,
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(f.config.ExpireDuration)),
		},
	}
}

func (f *ClaimFactory) GenerateKey(claims *jwt.RegisteredClaims) (string, error) {
	return f.config.GenerateKey(claims)
}

func (f *ClaimFactory) VerifyJwt(jwtString string, claim *jwt.RegisteredClaims) (bool, error) {
	return f.config.VerifyJwt(jwtString, claim)
}
