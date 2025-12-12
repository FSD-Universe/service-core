// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

//go:build http

// Package middleware
package middleware

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"half-nothing.cn/service-core/interfaces/http/dto"
	httpjwt "half-nothing.cn/service-core/interfaces/http/jwt"
)

var (
	ErrMissingOrMalformedJwt = errors.New("missing or malformed jwt")
)

func GetJWTMiddleware(factory httpjwt.ClaimFactoryInterface) echo.MiddlewareFunc {
	jwtConfig := echojwt.Config{
		TokenLookup: "header:Authorization:Bearer ",
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			if auth == "" {
				return nil, ErrMissingOrMalformedJwt
			}
			token := factory.EmptyClaim()
			_, err := factory.VerifyJwt(auth, token)
			if err != nil {
				return nil, err
			}
			return token, nil
		},
		ErrorHandler: func(c echo.Context, err error) error {
			var data *dto.ApiResponse[any]
			switch {
			case errors.Is(err, jwt.ErrTokenMalformed) || errors.Is(err, ErrMissingOrMalformedJwt):
				data = dto.NewApiResponse[any](dto.ErrMissingOrMalformedJwt, nil)
			case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenInvalidClaims):
				data = dto.NewApiResponse[any](dto.ErrInvalidOrExpiredJwt, nil)
			default:
				data = dto.NewApiResponse[any](dto.ErrUnknownJwtError, nil)
			}
			return data.Response(c)
		},
	}
	return echojwt.WithConfig(jwtConfig)
}
