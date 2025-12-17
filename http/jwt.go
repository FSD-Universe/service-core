// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

//go:build http && httpjwt

// Package http
package http

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

func jwtVerifyMiddleWare(flushToken bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			token := ctx.Get("user").(*jwt.Token)
			claim := token.Claims.(*httpjwt.Claims)
			if flushToken == claim.FlushToken {
				return next(ctx)
			}
			return dto.NewApiResponse[any](dto.ErrInvalidJwtType, nil).Response(ctx)
		}
	}
}

func GetJWTMiddleware(
	factory httpjwt.ClaimFactoryInterface,
) (
	jwtMiddleware echo.MiddlewareFunc,
	requireNoRefreshToken echo.MiddlewareFunc,
	requireRefreshToken echo.MiddlewareFunc,
) {
	jwtConfig := echojwt.Config{}
	jwtConfig.TokenLookup = "header:Authorization:Bearer "
	jwtConfig.ParseTokenFunc = func(c echo.Context, auth string) (interface{}, error) {
		if auth == "" {
			return nil, ErrMissingOrMalformedJwt
		}
		token := factory.EmptyClaim()
		_, err := factory.VerifyJwt(auth, token)
		if err != nil {
			return nil, err
		}
		return token, nil
	}
	jwtConfig.ErrorHandler = func(c echo.Context, err error) error {
		var data *dto.ApiResponse[any]
		var e *echojwt.TokenExtractionError
		if errors.As(err, &e) {
			err = e.Unwrap()
		}
		switch {
		case errors.Is(err, jwt.ErrTokenMalformed):
			fallthrough
		case errors.Is(err, ErrMissingOrMalformedJwt):
			fallthrough
		case err.Error() == "missing value in request header":
			data = dto.NewApiResponse[any](dto.ErrMissingOrMalformedJwt, nil)

		case errors.Is(err, jwt.ErrTokenExpired):
			fallthrough
		case errors.Is(err, jwt.ErrTokenInvalidClaims):
			data = dto.NewApiResponse[any](dto.ErrInvalidOrExpiredJwt, nil)

		default:
			data = dto.NewApiResponse[any](dto.ErrUnknownJwtError, nil)
		}
		return data.Response(c)
	}
	return echojwt.WithConfig(jwtConfig), jwtVerifyMiddleWare(false), jwtVerifyMiddleWare(true)
}
