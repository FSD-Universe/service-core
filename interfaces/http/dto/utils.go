// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

//go:build http

package dto

import (
	"github.com/labstack/echo/v4"
)

var (
	ErrLackParam             = NewApiStatus("PARAM_MISS", "缺少参数", HttpCodeBadRequest)
	ErrErrorParam            = NewApiStatus("PARAM_ERROR", "参数错误", HttpCodeBadRequest)
	ErrNoPermission          = NewApiStatus("NO_PERMISSION", "无权这么做", HttpCodePermissionDenied)
	ErrServerError           = NewApiStatus("SERVER_ERROR", "服务器错误", HttpCodeInternalError)
	ErrMissingOrMalformedJwt = NewApiStatus("MISSING_OR_MALFORMED_JWT", "缺少JWT令牌或者令牌格式错误", HttpCodeBadRequest)
	ErrInvalidOrExpiredJwt   = NewApiStatus("INVALID_OR_EXPIRED_JWT", "无效或过期的JWT令牌", HttpCodeUnauthorized)
	ErrInvalidJwtType        = NewApiStatus("INVALID_JWT_TYPE", "非法的JWT令牌类型", HttpCodeUnauthorized)
	ErrUnknownJwtError       = NewApiStatus("UNKNOWN_JWT_ERROR", "未知的JWT解析错误", HttpCodeInternalError)
	ErrRateLimitExceeded     = NewApiStatus("RATE_LIMIT_EXCEEDED", "请求频率过高", HttpCodeTooManyRequests)
	ErrNoMatchRoute          = NewApiStatus("NO_MATCH_ROUTE", "未匹配到路由", HttpCodeNotFound)
	ErrDataConflict          = NewApiStatus("DATA_CONFLICT", "数据冲突", HttpCodeConflict)
	SuccessHandleRequest     = NewApiStatus("SUCCESS", "成功", HttpCodeOk)
)

func ErrorResponse(ctx echo.Context, codeStatus *ApiStatus) error {
	return NewApiResponse[any](codeStatus, nil).Response(ctx)
}

func TextResponse(ctx echo.Context, httpCode int, content string) error {
	return ctx.String(httpCode, content)
}
