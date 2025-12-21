// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package utils
package utils

import "golang.org/x/crypto/bcrypt"

func BcryptEncrypt(data []byte, level int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(data, level)
}

func BcryptCompare(data []byte, hash []byte) bool {
	return bcrypt.CompareHashAndPassword(hash, data) == nil
}
