// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package jwt
package jwt

import (
	"testing"

	"half-nothing.cn/service-core/interfaces/config"
	"half-nothing.cn/service-core/interfaces/database/entity"
	"half-nothing.cn/service-core/testutils"
)

func TestClaimFactory(t *testing.T) {
	lg := testutils.NewFakeLogger(t)
	c := &config.JwtConfig{}
	c.InitDefaults()
	ok, err := c.Verify()
	if !ok {
		lg.Errorf("config verify error: %v", err)
		return
	}
	lg.Infof("config: %+v", c)
	factory := NewClaimFactory(c)
	claim := factory.CreateClaim(&entity.User{
		ID:         1,
		Cid:        1,
		Username:   "test",
		Permission: 1,
		Rating:     1,
	}, false)
	lg.Infof("claim: %+v", claim)
	token, err := factory.GenerateKey(claim)
	if err != nil {
		lg.Errorf("generate key error: %v", err)
		return
	}
	lg.Infof("token: %s", token)
	ok, err = factory.VerifyJwt(token, claim)
	decodeClaim := factory.EmptyClaim()
	ok, err = factory.VerifyJwt(token, decodeClaim)
	if err != nil {
		lg.Errorf("verify jwt error: %v", err)
		return
	}
	lg.Infof("verify jwt: %v", ok)
	lg.Infof("decode claim: %+v", decodeClaim)
}
