// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package utils
package utils

import (
	"fmt"
	"testing"
)

func TestEnum(t *testing.T) {
	var One = NewEnum[int, string](1, "One")
	var Two = NewEnum[int, string](2, "Two")
	var Three = NewEnum[int, string](3, "Three")
	var Four = NewEnum[int, string](4, "Four")
	var Five = NewEnum[int, string](5, "Five")
	var numbers = NewEnums(One, Two, Three, Four, Five)

	if !numbers.IsValidEnum(1) {
		t.Errorf("Expected true, got false")
	}

	if numbers.IsValidEnum(6) {
		t.Errorf("Expected false, got true")
	}

	if numbers.GetEnum(1) != One {
		t.Errorf("Expected %v, got %v", One, numbers.GetEnum(1))
	}

	if numbers.GetEnum(6) != nil {
		t.Errorf("Expected nil, got %v", numbers.GetEnum(6))
	}
}

func ExampleNewEnum() {
	var One = NewEnum[int, string](1, "One")
	var Two = NewEnum[int, string](2, "Two")
	fmt.Println(One.Value)
	fmt.Println(Two.Data)
	fmt.Println(One.Value == 1)
	fmt.Println(Two.Data == "One")
	// output:
	// 1
	// Two
	// true
	// false
}

func ExampleNewEnums() {
	var One = NewEnum[int, string](1, "One")
	var Two = NewEnum[int, string](2, "Two")
	var Three = NewEnum[int, string](3, "Three")
	var numbers = NewEnums(One, Two, Three)
	fmt.Println(numbers.IsValidEnum(1))
	fmt.Println(numbers.IsValidEnum(4))
	fmt.Println(numbers.GetEnum(1) == One)
	fmt.Println(numbers.GetEnum(4) == nil)
	// output:
	// true
	// false
	// true
	// true
}
