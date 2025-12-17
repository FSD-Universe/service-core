// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package utils
package utils

import "testing"

func TestGetOrDefault(t *testing.T) {
	type args[T any] struct {
		data        T
		defaultData T
	}
	tests := []struct {
		name string
		args args[int]
		want int
	}{
		{name: "TestGetOrDefault", args: args[int]{data: 0, defaultData: 1}, want: 1},
		{name: "TestGetOrDefault", args: args[int]{data: 1, defaultData: 2}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOrDefault(tt.args.data, tt.args.defaultData); got != tt.want {
				t.Errorf("GetOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
