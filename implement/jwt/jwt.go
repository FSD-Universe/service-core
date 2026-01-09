// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

// Package jwt
package jwt

import (
	"time"

	c "half-nothing.cn/service-core/interfaces/config"
	"half-nothing.cn/service-core/interfaces/database/entity"
)

type ClaimFactory struct {
	*c.JwtConfig
}

func NewClaimFactory(config *c.JwtConfig) *ClaimFactory {
	return &ClaimFactory{
		JwtConfig: config,
	}
}

func (f *ClaimFactory) EmptyClaim() *Claims {
	return &Claims{}
}

func (f *ClaimFactory) EmptyFsdClaim() *FsdClaims {
	return &FsdClaims{}
}

func (f *ClaimFactory) CreateClaim(user *entity.User, flushToken bool) *Claims {
	var expiredDuration time.Duration
	if flushToken {
		expiredDuration = f.JwtConfig.RefreshExpireDuration
	} else {
		expiredDuration = f.JwtConfig.ExpireDuration
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
			Issuer:    f.JwtConfig.Issuer,
			Subject:   user.Username,
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(expiredDuration)),
		},
	}
}

func (f *ClaimFactory) CreateFsdClaim(user *entity.User) *FsdClaims {
	now := time.Now()
	return &FsdClaims{
		ControllerRating: user.Rating,
		PilotRating:      user.Rating,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    f.JwtConfig.Issuer,
			Subject:   user.Username,
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(f.JwtConfig.ExpireDuration)),
		},
	}
}

func (f *ClaimFactory) GetJWTConfig() *c.JwtConfig {
	return f.JwtConfig
}
