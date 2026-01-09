// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

// Package utils
package utils

import "testing"

func TestBcrypt(t *testing.T) {
	type args struct {
		data  []byte
		level int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "TestBcryptEncrypt", args: args{data: []byte("123456"), level: 10}},
		{name: "TestBcryptEncrypt", args: args{data: []byte("123456"), level: 5}},
		{name: "TestBcryptEncrypt", args: args{data: []byte("123456"), level: 12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := BcryptEncrypt(tt.args.data, tt.args.level)
			if err != nil {
				t.Errorf("BcryptEncrypt() error = %v", err)
				return
			}
			if !BcryptCompare(tt.args.data, hash) {
				t.Errorf("BcryptCompare() = %v, want %v", false, true)
			}
		})
	}
}
