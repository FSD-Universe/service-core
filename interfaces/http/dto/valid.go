// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

// Package dto
package dto

import (
	"fmt"
	"reflect"
	"regexp"
	"slices"
	"strings"

	"half-nothing.cn/service-core/utils"
)

type TagHandler func(field reflect.Value, tagValue string, additional []string) (*ApiStatus, error)

const (
	TagName = "valid"

	TagRequired = "required"
	TagMin      = "min"
	TagMax      = "max"
	TagRegex    = "regex"
	TagLength   = "length"

	AdditionalExclude = "exclude"
)

var tagHandlers = map[string]TagHandler{
	TagMin:    processMin,
	TagMax:    processMax,
	TagRegex:  processRegex,
	TagLength: processLength,
}

func isStruct(v reflect.Value) bool {
	if v.Kind() == reflect.Struct {
		return true
	}
	for v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return false
		}
		v = v.Elem()
	}
	return v.Kind() == reflect.Struct
}

func isArrayOrSlice(v reflect.Value) bool {
	if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
		return true
	}
	for v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return false
		}
		v = v.Elem()
	}
	return v.Kind() == reflect.Slice || v.Kind() == reflect.Array
}

func getArrayOrSliceType(v reflect.Value) reflect.Kind {
	types := v.Type()
	for types.Kind() == reflect.Pointer {
		types = types.Elem()
	}
	subType := types.Elem()
	for subType.Kind() == reflect.Pointer {
		subType = subType.Elem()
	}
	return subType.Kind()
}

func ValidStruct(val interface{}) (res *ApiStatus, err error) {
	v := reflect.ValueOf(val)
	for v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return nil, nil
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, nil
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := v.Type().Field(i).Tag.Get(TagName)
		required := strings.Contains(tag, TagRequired)

		if required && field.IsZero() {
			return ErrLackParam, nil
		}

		if field.IsZero() {
			continue
		}

		if isArrayOrSlice(field) && getArrayOrSliceType(field) == reflect.Struct {
			for field.Kind() == reflect.Pointer {
				field = field.Elem()
			}
			for i := 0; i < field.Len(); i++ {
				res, err = ValidStruct(field.Index(i).Interface())
				if res != nil || err != nil {
					return
				}
			}
			continue
		}

		if isStruct(field) {
			res, err = ValidStruct(field.Interface())
			if res != nil || err != nil {
				return
			}
			continue
		}

		if tag == "" {
			continue
		}

		tags := strings.Split(tag, ",")

		for field.Kind() == reflect.Pointer {
			field = field.Elem()
		}

		for _, t := range tags {
			if t == TagRequired {
				continue
			}
			tags := strings.SplitN(t, "=", 2)
			tagName := tags[0] // min max regex length
			if len(tags) != 2 {
				return nil, fmt.Errorf("tag %s error, miss argument", tagName)
			}
			tags = strings.Split(tags[1], ";")
			tagValue := tags[0]    // 0 128 ^[A-Za-z_-][\\w-]*$
			additional := tags[1:] // exclude
			handler, ok := tagHandlers[tagName]
			if !ok {
				return nil, fmt.Errorf("tag %s unsupport", tagName)
			}
			res, err = handler(field, tagValue, additional)
			if res != nil || err != nil {
				return
			}
		}
	}

	return nil, nil
}

func processLength(field reflect.Value, tagValue string, _ []string) (*ApiStatus, error) {
	if field.Kind() != reflect.String {
		return nil, fmt.Errorf("tag 'length' unsupport type: %v", field.Kind())
	}
	target := utils.StrToInt(tagValue, -1)
	if target == -1 {
		return nil, fmt.Errorf("tag 'length' error, illegal argument %s", tagValue)
	}
	switch field.Kind() {
	case reflect.String:
		if field.Len() != target {
			return ErrErrorParam, nil
		}
	case reflect.Slice:
		if field.Len() != target {
			return ErrErrorParam, nil
		}
	default:
		return nil, fmt.Errorf("tag 'length' unsupport type: %v", field.Kind())
	}
	return nil, nil
}

func processRegex(field reflect.Value, tagValue string, _ []string) (*ApiStatus, error) {
	ok, err := regexp.MatchString(tagValue, field.String())
	if err != nil {
		return nil, err
	}
	if !ok {
		return ErrErrorParam, nil
	}
	return nil, nil
}

func processMax(field reflect.Value, tagValue string, additional []string) (*ApiStatus, error) {
	target := utils.StrToFloat(tagValue, -1)
	if target == -1 {
		return nil, fmt.Errorf("tag 'max' error, illegal argument %s", tagValue)
	}
	if slices.Contains(additional, AdditionalExclude) {
		return checkMaxExclude(field, target)
	}
	return checkMax(field, target)
}

func processMin(field reflect.Value, tagValue string, additional []string) (*ApiStatus, error) {
	target := utils.StrToFloat(tagValue, -1)
	if target == -1 {
		return nil, fmt.Errorf("tag 'min' error, illegal argument %s", tagValue)
	}
	if slices.Contains(additional, AdditionalExclude) {
		return checkMinExclude(field, target)
	}
	return checkMin(field, target)
}

func checkMinExclude(val reflect.Value, target float64) (*ApiStatus, error) {
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if val.Int() <= int64(target) {
			return ErrErrorParam, nil
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if val.Uint() <= uint64(target) {
			return ErrErrorParam, nil
		}
	case reflect.Float32, reflect.Float64:
		if val.Float() <= target {
			return ErrErrorParam, nil
		}
	case reflect.String:
		if val.Len() <= int(target) {
			return ErrErrorParam, nil
		}
	case reflect.Pointer:
		return checkMin(val.Elem(), target)
	case reflect.Slice:
		fallthrough
	case reflect.Array:
		for i := 0; i < val.Len(); i++ {
			res, err := checkMin(val.Index(i), target)
			if res != nil || err != nil {
				return res, err
			}
		}
	default:
		return nil, fmt.Errorf("tag 'min' unsupport type: %v", val.Kind())
	}
	return nil, nil
}

func checkMin(val reflect.Value, target float64) (*ApiStatus, error) {
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if val.Int() < int64(target) {
			return ErrErrorParam, nil
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if val.Uint() < uint64(target) {
			return ErrErrorParam, nil
		}
	case reflect.Float32, reflect.Float64:
		if val.Float() < target {
			return ErrErrorParam, nil
		}
	case reflect.String:
		if val.Len() < int(target) {
			return ErrErrorParam, nil
		}
	case reflect.Pointer:
		return checkMin(val.Elem(), target)
	case reflect.Slice:
		fallthrough
	case reflect.Array:
		for i := 0; i < val.Len(); i++ {
			res, err := checkMin(val.Index(i), target)
			if res != nil || err != nil {
				return res, err
			}
		}
	default:
		return nil, fmt.Errorf("tag 'min' unsupport type: %v", val.Kind())
	}
	return nil, nil
}

func checkMaxExclude(val reflect.Value, target float64) (*ApiStatus, error) {
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if val.Int() >= int64(target) {
			return ErrErrorParam, nil
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if val.Uint() >= uint64(target) {
			return ErrErrorParam, nil
		}
	case reflect.Float32, reflect.Float64:
		if val.Float() >= target {
			return ErrErrorParam, nil
		}
	case reflect.String:
		if val.Len() >= int(target) {
			return ErrErrorParam, nil
		}
	case reflect.Pointer:
		return checkMax(val.Elem(), target)
	case reflect.Slice:
		fallthrough
	case reflect.Array:
		for i := 0; i < val.Len(); i++ {
			res, err := checkMax(val.Index(i), target)
			if res != nil || err != nil {
				return res, err
			}
		}
	default:
		return nil, fmt.Errorf("tag 'max' unsupport type: %v", val.Kind())
	}
	return nil, nil
}

func checkMax(val reflect.Value, target float64) (*ApiStatus, error) {
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if val.Int() > int64(target) {
			return ErrErrorParam, nil
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if val.Uint() > uint64(target) {
			return ErrErrorParam, nil
		}
	case reflect.Float32, reflect.Float64:
		if val.Float() > target {
			return ErrErrorParam, nil
		}
	case reflect.String:
		if val.Len() > int(target) {
			return ErrErrorParam, nil
		}
	case reflect.Pointer:
		return checkMax(val.Elem(), target)
	case reflect.Slice:
		fallthrough
	case reflect.Array:
		for i := 0; i < val.Len(); i++ {
			res, err := checkMax(val.Index(i), target)
			if res != nil || err != nil {
				return res, err
			}
		}
	default:
		return nil, fmt.Errorf("tag 'max' unsupport type: %v", val.Kind())
	}
	return nil, nil
}
