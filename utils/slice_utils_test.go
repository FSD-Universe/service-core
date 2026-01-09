// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

// Package utils
package utils

import (
	"reflect"
	"testing"
)

func TestAny(t *testing.T) {
	type args struct {
		src        []int
		comparator func(element int) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Any",
			args: args{
				src: []int{1, 2, 3, 4, 5},
				comparator: func(element int) bool {
					return element > 3
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Any(tt.args.src, tt.args.comparator); got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		src    []int
		filter func(element int) bool
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{
			name: "Filter",
			args: args{
				src: []int{1, 2, 3, 4, 5},
				filter: func(element int) bool {
					return element > 3
				},
			},
			wantResult: []int{4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Filter(tt.args.src, tt.args.filter); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Filter() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestFind(t *testing.T) {
	type args struct {
		src        []int
		comparator func(element int) bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Find",
			args: args{
				src: []int{1, 2, 3, 4, 5},
				comparator: func(element int) bool {
					return element > 3
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.args.src, tt.args.comparator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForEach(t *testing.T) {
	type args struct {
		src      []int
		callback func(index int, element int)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ForEach",
			args: args{
				src: []int{1, 2, 3, 4, 5},
				callback: func(index int, element int) {
					t.Log(index, element)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ForEach(tt.args.src, tt.args.callback)
		})
	}
}

func TestMap(t *testing.T) {
	type args struct {
		src    []int
		mapper func(element int) int
	}
	type testCase struct {
		name string
		args args
		want []int
	}
	tests := []testCase{
		{
			name: "Map",
			args: args{
				src: []int{1, 2, 3, 4, 5},
				mapper: func(element int) int {
					return element * 2
				},
			},
			want: []int{2, 4, 6, 8, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Map(tt.args.src, tt.args.mapper)
			if !reflect.DeepEqual(tt.args.src, tt.want) {
				t.Errorf("Map() = %v, want %v", tt.args.src, tt.want)
			}
		})
	}
}

func TestReverseForEach(t *testing.T) {
	type args struct {
		slice []int
		f     func(index int, value int)
	}
	type testCase struct {
		name string
		args args
	}
	tests := []testCase{
		{
			name: "ReverseForEach",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				f: func(index int, value int) {
					t.Log(index, value)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseForEach(tt.args.slice, tt.args.f)
		})
	}
}
