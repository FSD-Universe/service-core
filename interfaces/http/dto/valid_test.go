// Package dto
package dto

import (
	"errors"
	"testing"
)

func TestValidStruct(t *testing.T) {
	type arg struct {
		Name string `valid:"required,min=2,max=10"`
		Age  int    `valid:"required,min=18,max=30"`
		Test string `valid:"required,regex=^[A-Za-z_-][\\w-]*$"`
		User string `valid:"length=6"`
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
		{name: "TestValidStruct", args: &arg{Name: "aaa", Age: 23, Test: "asd", User: "AAAAA"}, want: &want{want: nil, hasError: false}},
		{name: "TestValidStruct", args: &arg{Name: "aaa", Age: 23, Test: "asd", User: "AAAAAA"}, want: &want{want: nil, hasError: false}},
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
