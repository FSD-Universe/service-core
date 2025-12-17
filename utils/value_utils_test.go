// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package utils
package utils

import "testing"

func TestGetPointerDataOrDefault(t *testing.T) {
	a := 4
	type args struct {
		data        *int
		defaultData int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "TestGetPointerDataOrDefault", args: args{data: nil, defaultData: 3}, want: 3},
		{name: "TestGetPointerDataOrDefault", args: args{data: &a, defaultData: 0}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPointerDataOrDefault(tt.args.data, tt.args.defaultData); got != tt.want {
				t.Errorf("GetPointerDataOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPointerData(t *testing.T) {
	a := 5
	type args struct {
		data *int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "TestGetPointerData", args: args{data: &a}, want: 5},
		{name: "TestGetPointerData", args: args{data: nil}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPointerData(tt.args.data); got != tt.want {
				t.Errorf("GetPointerData() = %v, want %v", got, tt.want)
			}
		})
	}
}
