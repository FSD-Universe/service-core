//go:build http && httpjwt

// Package config
package config

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/thanhpk/randstr"
	"half-nothing.cn/service-core/utils"
)

type TokenFunc = func() any

type JwtConfig struct {
	Issuer        string `yaml:"issuer"`
	SignMethod    string `yaml:"sign_method"`
	Secret        string `yaml:"secret"`
	PublicKey     string `yaml:"public_key"`
	PrivateKey    string `yaml:"private_key"`
	Expire        string `yaml:"expire"`
	RefreshExpire string `yaml:"refresh_expire"`

	// 内部变量
	SecretContent         string          `yaml:"-"`
	PublicKeyContent      *rsa.PublicKey  `yaml:"-"`
	PrivateKeyContent     *rsa.PrivateKey `yaml:"-"`
	JWTSignMethod         SignMethod      `yaml:"-"`
	JWTTokenFunc          TokenFunc       `yaml:"-"`
	ExpireDuration        time.Duration   `yaml:"-"`
	RefreshExpireDuration time.Duration   `yaml:"-"`
}

func (j *JwtConfig) InitDefaults() {
	j.Issuer = "service"
	j.SignMethod = "HS256"
	j.Secret = ""
	j.PublicKey = "public.pem"
	j.PrivateKey = "private.pem"
	j.Expire = "1h"
	j.RefreshExpire = "24h"
}

type SignMethod *utils.Enum[string, jwt.SigningMethod]

var (
	SignMethodHS256 = utils.NewEnum[string, jwt.SigningMethod]("HS256", jwt.SigningMethodHS256)
	SignMethodHS384 = utils.NewEnum[string, jwt.SigningMethod]("HS384", jwt.SigningMethodHS384)
	SignMethodHS512 = utils.NewEnum[string, jwt.SigningMethod]("HS512", jwt.SigningMethodHS512)
	SignMethodRS256 = utils.NewEnum[string, jwt.SigningMethod]("RS256", jwt.SigningMethodRS256)
	SignMethodRS384 = utils.NewEnum[string, jwt.SigningMethod]("RS384", jwt.SigningMethodRS384)
	SignMethodRS512 = utils.NewEnum[string, jwt.SigningMethod]("RS512", jwt.SigningMethodRS512)
)

var signMethods = utils.NewEnums(
	SignMethodHS256,
	SignMethodHS384,
	SignMethodHS512,
	SignMethodRS256,
	SignMethodRS384,
	SignMethodRS512,
)

func (j *JwtConfig) HMACToken() any {
	return []byte(j.SecretContent)
}

func (j *JwtConfig) RSAToken() any {
	return j.PrivateKeyContent
}

func (j *JwtConfig) GenerateKey(claims *jwt.RegisteredClaims) (string, error) {
	return jwt.NewWithClaims(j.JWTSignMethod.Data, claims).SignedString(j.JWTTokenFunc())
}

func (j *JwtConfig) defaultKeyFunc(token *jwt.Token) (interface{}, error) {
	if token.Method.Alg() != j.SignMethod {
		return nil, errors.New("illegal signature methods")
	}
	return j.JWTTokenFunc(), nil
}

func (j *JwtConfig) VerifyJwt(jwtString string, claim *jwt.RegisteredClaims) (bool, error) {
	parseClaim, err := jwt.ParseWithClaims(jwtString, claim, j.defaultKeyFunc)
	if err != nil {
		return false, err
	}
	claim, ok := parseClaim.Claims.(*jwt.RegisteredClaims)
	return ok, nil
}

func (j *JwtConfig) Verify() (bool, error) {
	if duration, err := time.ParseDuration(j.Expire); err == nil {
		j.ExpireDuration = duration
	} else {
		return false, fmt.Errorf("expire is not a valid duration: %v", err)
	}
	if duration, err := time.ParseDuration(j.RefreshExpire); err == nil {
		j.RefreshExpireDuration = duration
	} else {
		return false, fmt.Errorf("refresh expire is not a valid duration: %v", err)
	}
	if j.RefreshExpireDuration < j.ExpireDuration {
		return false, fmt.Errorf("refresh expire must be greater than expire")
	}
	return j.verifySignMethod()
}

func (j *JwtConfig) verifySignMethod() (bool, error) {
	if j.SignMethod == "" {
		return false, fmt.Errorf("sign method is empty")
	}
	if !signMethods.IsValidEnum(j.SignMethod) {
		return false, fmt.Errorf("sign method is not allowed")
	}
	j.JWTSignMethod = signMethods.GetEnum(j.SignMethod)
	switch j.JWTSignMethod {
	case SignMethodHS256:
	case SignMethodHS384:
	case SignMethodHS512:
		if j.Secret != "" {
			j.SecretContent = j.Secret
		} else {
			j.SecretContent = randstr.String(64)
		}
	case SignMethodRS256:
	case SignMethodRS384:
	case SignMethodRS512:
		if j.PrivateKey == "" {
			return false, fmt.Errorf("private key is empty")
		}
		if j.PublicKey == "" {
			return false, fmt.Errorf("public key is empty")
		}

		publicKey, err := os.ReadFile(j.PublicKey)
		if err != nil {
			return false, fmt.Errorf("read public key failed: %v", err)
		}
		j.PublicKeyContent, err = jwt.ParseRSAPublicKeyFromPEM(publicKey)
		if err != nil {
			return false, fmt.Errorf("parse public key failed: %v", err)
		}

		privateKey, err := os.ReadFile(j.PrivateKey)
		if err != nil {
			return false, fmt.Errorf("read private key failed: %v", err)
		}
		j.PrivateKeyContent, err = jwt.ParseRSAPrivateKeyFromPEM(privateKey)
		if err != nil {
			return false, fmt.Errorf("parse private key failed: %v", err)
		}
	}
	return true, nil
}
