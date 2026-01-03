// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

//go:build database

// Package dto
package dto

var (
	ErrDatabaseError  = NewApiStatus("DATABASE_ERROR", "数据库错误", HttpCodeInternalError)
	ErrRecordNotFound = NewApiStatus("RECORD_NOT_FOUND", "指定记录未找到", HttpCodeNotFound)
	ErrDataConflict   = NewApiStatus("DATA_CONFLICT", "数据冲突", HttpCodeConflict)
)
