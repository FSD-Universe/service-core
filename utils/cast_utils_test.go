// Package utils
// $END$
package utils

import (
	"errors"
	"testing"
	"time"
)

func TestStrToFloat(t *testing.T) {
	type args struct {
		str          string
		defaultValue float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "TestStrToFloat", args: args{str: "123.456", defaultValue: 0.0}, want: 123.456},
		{name: "TestStrToFloat", args: args{str: "abc", defaultValue: 0.0}, want: 0.0},
		{name: "TestStrToFloat", args: args{str: "", defaultValue: 0.0}, want: 0.0},
		{name: "TestStrToFloat", args: args{str: "123", defaultValue: 0.0}, want: 123.0},
		{name: "TestStrToFloat", args: args{str: "123.456.789", defaultValue: 0.0}, want: 0.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrToFloat(tt.args.str, tt.args.defaultValue); got != tt.want {
				t.Errorf("StrToFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrToInt(t *testing.T) {
	type args struct {
		str          string
		defaultValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "TestStrToInt", args: args{str: "123", defaultValue: 0}, want: 123},
		{name: "TestStrToInt", args: args{str: "abc", defaultValue: 0}, want: 0},
		{name: "TestStrToInt", args: args{str: "", defaultValue: 0}, want: 0},
		{name: "TestStrToInt", args: args{str: "123.456", defaultValue: 0}, want: 0},
		{name: "TestStrToInt", args: args{str: "123.456.789", defaultValue: 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrToInt(tt.args.str, tt.args.defaultValue); got != tt.want {
				t.Errorf("StrToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseDuration(t *testing.T) {
	type args struct {
		duration string
	}
	tests := []struct {
		name string
		args args
		want time.Duration
		err  error
	}{
		{name: "TestParseDuration", args: args{duration: "1h"}, want: time.Hour, err: nil},
		{name: "TestParseDuration", args: args{duration: "1h20m"}, want: time.Hour + 20*time.Minute, err: nil},
		{name: "TestParseDuration", args: args{duration: "1d"}, want: 0, err: ErrInvalidDuration},
		{name: "TestParseDuration", args: args{duration: "abc"}, want: 0, err: ErrInvalidDuration},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			target := time.Duration(0)
			err := ParseDuration(tt.args.duration, &target)
			if !errors.Is(err, tt.err) {
				t.Errorf("ParseDuration() error = %v, want %v", err, tt.err)
			}
			if target != tt.want {
				t.Errorf("ParseDuration() = %v, want %v", target, tt.want)
			}
		})
	}
}

func BenchmarkParseDuration(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		target := time.Duration(0)
		_ = ParseDuration("1h20m", &target)
	}
}

func BenchmarkStrToInt(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StrToInt("123", 0)
	}
}

func BenchmarkStrToFloat(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StrToFloat("123.456", 0.0)
	}
}
