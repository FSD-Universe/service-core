// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

// Package dto
package dto

import (
	"errors"
	"testing"
)

//goland:noinspection DuplicatedCode
func TestValidStruct(t *testing.T) {
	type arg struct {
		Name  string `valid:"required,min=2,max=10"`
		Age   int    `valid:"required,min=18,max=30"`
		Test  string `valid:"required,regex=^[A-Za-z_-][\\w-]*$"`
		User  string `valid:"length=6"`
		Home  string `valid:"min=10,max=12;exclude"`
		Phone string `valid:"min=10;exclude,max=12"`
	}
	type want struct {
		want     *ApiStatus
		hasError bool
	}
	tests := []struct {
		name string
		args *arg
		want *want
	}{
		{name: "TestValidStruct", args: &arg{Name: "AAA", Age: 0, Test: "AAA"}, want: &want{want: ErrLackParam, hasError: false}},
		{name: "TestValidStruct", args: &arg{Name: "aaa", Age: 24, Test: "1AAAA"}, want: &want{want: ErrErrorParam, hasError: false}},
		{name: "TestValidStruct", args: &arg{Name: "aaa", Age: 23, Test: "asd"}, want: &want{want: nil, hasError: false}},
		{name: "TestValidStruct", args: &arg{Name: "aaa", Age: 23, Test: "asd", User: "AAAAA"}, want: &want{want: ErrErrorParam, hasError: false}},
		{name: "TestValidStruct", args: &arg{Name: "aaa", Age: 23, Test: "asd", User: "AAAAAA"}, want: &want{want: nil, hasError: false}},
		{name: "TestValidStruct", args: &arg{Name: "aa", Age: 18, Test: "asd"}, want: &want{want: nil, hasError: false}},
		{name: "TestValidStruct", args: &arg{Name: "aaaaaaaaaa", Age: 30, Test: "asd"}, want: &want{want: nil, hasError: false}},
		{name: "TestValidStruct", args: &arg{Name: "aaa", Age: 18, Test: "asd", Home: "AAAAAAAAAAAA"}, want: &want{want: ErrErrorParam, hasError: false}},
		{name: "TestValidStruct", args: &arg{Name: "aaa", Age: 18, Test: "asd", Phone: "AAAAAAAAAA"}, want: &want{want: ErrErrorParam, hasError: false}},
		{name: "TestValidStruct", args: &arg{Name: "aaa", Age: 18, Test: "asd"}, want: &want{want: nil, hasError: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidStruct(tt.args)
			if (err != nil) != tt.want.hasError {
				t.Errorf("ValidStruct() error = %v, wantErr %v", err, tt.want.hasError)
				return
			}
			if !tt.want.hasError {
				if !errors.Is(got, tt.want.want) {
					t.Errorf("ValidStruct() got = %v, want %v", got, tt.want.want)
				}
			}
		})
	}
}

//goland:noinspection DuplicatedCode
func TestNestStruct(t *testing.T) {
	type TestStruct struct {
		Name string `valid:"required,min=2,max=10"`
	}
	type arg struct {
		Test  TestStruct `valid:"required"`
		Test2 *TestStruct
	}
	type want struct {
		want     *ApiStatus
		hasError bool
	}
	tests := []struct {
		name string
		args *arg
		want *want
	}{
		{name: "TestNestStruct", args: &arg{Test: TestStruct{}}, want: &want{want: ErrLackParam, hasError: false}},
		{name: "TestNestStruct", args: &arg{Test: TestStruct{Name: "TEST"}, Test2: &TestStruct{}}, want: &want{want: ErrLackParam, hasError: false}},
		{name: "TestNestStruct", args: &arg{Test: TestStruct{Name: "TEST"}, Test2: &TestStruct{Name: "TEST"}}, want: &want{want: nil, hasError: false}},
		{name: "TestNestStruct", args: &arg{Test: TestStruct{Name: "TEST"}}, want: &want{want: nil, hasError: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidStruct(tt.args)
			if (err != nil) != tt.want.hasError {
				t.Errorf("ValidStruct() error = %v, wantErr %v", err, tt.want.hasError)
				return
			}
			if !tt.want.hasError {
				if !errors.Is(got, tt.want.want) {
					t.Errorf("ValidStruct() got = %v, want %v", got, tt.want.want)
				}
			}
		})
	}
}

//goland:noinspection DuplicatedCode
func TestNestArrayAndSlice(t *testing.T) {
	type TestStruct struct {
		Name string `valid:"required,min=2,max=10"`
	}
	type arg struct {
		Test  []TestStruct `valid:"required"`
		Test2 []*TestStruct
	}
	type want struct {
		want     *ApiStatus
		hasError bool
	}
	tests := []struct {
		name string
		args *arg
		want *want
	}{
		{name: "TestNestArrayAndSlice", args: &arg{}, want: &want{want: ErrLackParam, hasError: false}},
		{name: "TestNestArrayAndSlice", args: &arg{Test2: []*TestStruct{}}, want: &want{want: ErrLackParam, hasError: false}},
		{name: "TestNestArrayAndSlice", args: &arg{Test: []TestStruct{{Name: "TEST"}}, Test2: []*TestStruct{{Name: "TEST"}}}, want: &want{want: nil, hasError: false}},
		{name: "TestNestArrayAndSlice", args: &arg{Test: []TestStruct{{Name: "TEST"}}}, want: &want{want: nil, hasError: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidStruct(tt.args)
			if (err != nil) != tt.want.hasError {
				t.Errorf("ValidStruct() error = %v, wantErr %v", err, tt.want.hasError)
				return
			}
			if !tt.want.hasError {
				if !errors.Is(got, tt.want.want) {
					t.Errorf("ValidStruct() got = %v, want %v", got, tt.want.want)
				}
			}
		})
	}
}

//goland:noinspection DuplicatedCode
func TestNestStructAndArrayAndSlice(t *testing.T) {
	type TestStruct struct {
		Name string `valid:"required,min=2,max=10"`
	}
	testPtr := &TestStruct{Name: "T"}
	testPtrPtr := &testPtr
	type arg struct {
		Test  []TestStruct `valid:"required"`
		Test2 []*TestStruct
		Test3 TestStruct `valid:"required"`
		Test4 *TestStruct
		Test5 []**TestStruct
		Test6 **TestStruct
		Test7 **[]**TestStruct
	}
	testStruct := []**TestStruct{testPtrPtr}
	testStructPtr := &testStruct
	testStructPtrPtr := &testStructPtr
	type want struct {
		want     *ApiStatus
		hasError bool
	}
	tests := []struct {
		name string
		args *arg
		want *want
	}{
		{name: "TestNestStructAndArrayAndSlice", args: &arg{}, want: &want{want: ErrLackParam, hasError: false}},
		{name: "TestNestStructAndArrayAndSlice", args: &arg{Test2: []*TestStruct{}}, want: &want{want: ErrLackParam, hasError: false}},
		{name: "TestNestStructAndArrayAndSlice", args: &arg{Test: []TestStruct{{Name: "TEST"}}}, want: &want{want: ErrLackParam, hasError: false}},
		{name: "TestNestStructAndArrayAndSlice", args: &arg{Test: []TestStruct{{Name: "TEST"}}, Test3: TestStruct{Name: "TEST"}}, want: &want{want: nil, hasError: false}},
		{name: "TestNestStructAndArrayAndSlice", args: &arg{Test: []TestStruct{{Name: "TEST"}}, Test3: TestStruct{Name: "TEST"}, Test6: testPtrPtr}, want: &want{want: ErrErrorParam, hasError: false}},
		{name: "TestNestStructAndArrayAndSlice", args: &arg{Test: []TestStruct{{Name: "TEST"}}, Test3: TestStruct{Name: "TEST"}, Test5: []**TestStruct{testPtrPtr}}, want: &want{want: ErrErrorParam, hasError: false}},
		{name: "TestNestStructAndArrayAndSlice", args: &arg{Test: []TestStruct{{Name: "TEST"}}, Test3: TestStruct{Name: "TEST"}, Test7: testStructPtrPtr}, want: &want{want: ErrErrorParam, hasError: false}},
		{name: "TestNestStructAndArrayAndSlice", args: &arg{Test: []TestStruct{{Name: "T"}}, Test3: TestStruct{Name: "TEST"}}, want: &want{want: ErrErrorParam, hasError: false}},
		{name: "TestNestStructAndArrayAndSlice", args: &arg{Test: []TestStruct{{Name: "TEST"}}, Test3: TestStruct{Name: "TESTTESTTEST"}}, want: &want{want: ErrErrorParam, hasError: false}},
		{name: "TestNestStructAndArrayAndSlice", args: &arg{Test: []TestStruct{{Name: "TEST"}}, Test2: []*TestStruct{{Name: "T"}}, Test3: TestStruct{Name: "TEST"}, Test4: &TestStruct{Name: "TEST"}}, want: &want{want: ErrErrorParam, hasError: false}},
		{name: "TestNestStructAndArrayAndSlice", args: &arg{Test: []TestStruct{{Name: "TEST"}}, Test3: TestStruct{Name: "TEST"}, Test4: &TestStruct{Name: "T"}}, want: &want{want: ErrErrorParam, hasError: false}},
		{name: "TestNestStructAndArrayAndSlice", args: &arg{Test: []TestStruct{{Name: "TEST"}}, Test2: []*TestStruct{{Name: "TEST"}}, Test3: TestStruct{Name: "TEST"}, Test4: &TestStruct{Name: "TEST"}}, want: &want{want: nil, hasError: false}},
		{name: "TestNestStructAndArrayAndSlice", args: &arg{Test: []TestStruct{{Name: "TEST"}}, Test2: []*TestStruct{{Name: "TEST"}}, Test3: TestStruct{Name: "TEST"}}, want: &want{want: nil, hasError: false}},
		{name: "TestNestStructAndArrayAndSlice", args: &arg{Test: []TestStruct{{Name: "TEST"}}, Test3: TestStruct{Name: "TEST"}, Test4: &TestStruct{Name: "TEST"}}, want: &want{want: nil, hasError: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidStruct(tt.args)
			if (err != nil) != tt.want.hasError {
				t.Errorf("ValidStruct() error = %v, wantErr %v", err, tt.want.hasError)
				return
			}
			if !tt.want.hasError {
				if !errors.Is(got, tt.want.want) {
					t.Errorf("ValidStruct() got = %v, want %v", got, tt.want.want)
				}
			}
		})
	}
}

func BenchmarkValidStruct(b *testing.B) {
	type Country struct {
		Name string `valid:"required,min=2,max=10"`
		Code string `valid:"required,min=2,max=10"`
	}
	type Child struct {
		Name string `valid:"required,min=2,max=10"`
		Age  int    `valid:"required,min=18,max=30"`
	}
	type TestStruct struct {
		Name     string   `valid:"required,min=2,max=10"`
		Age      int      `valid:"required,min=18,max=30"`
		Username string   `valid:"required,regex=^[A-Za-z_-][\\w-]*$"`
		Country  Country  `valid:"required"`
		Children []*Child `valid:"required"`
	}
	test := &TestStruct{
		Name:     "TEST",
		Age:      20,
		Username: "TEST",
		Country:  Country{Name: "TEST", Code: "TEST"},
		Children: []*Child{{Name: "TEST1", Age: 20}, {Name: "TEST2", Age: 20}},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ValidStruct(test)
	}
}
