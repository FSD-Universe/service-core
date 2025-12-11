// Package dto
package dto

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"half-nothing.cn/service-core/utils"
)

func ValidStruct(val interface{}) (*ApiStatus, error) {
	v := reflect.ValueOf(val)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get("valid")
		if tag == "" {
			continue
		}
		tags := strings.Split(tag, ",")
		for _, t := range tags {
			if t == "required" && v.Field(i).IsZero() {
				return ErrLackParam, nil
			}
			if strings.HasPrefix(t, "min") {
				res, err := processTagMin(t, v.Field(i))
				if res == nil && err == nil {
					continue
				}
				return res, err
			}
			if strings.HasPrefix(t, "max") {
				res, err := processTagMax(t, v.Field(i))
				if res == nil && err == nil {
					continue
				}
				return res, err
			}
			if strings.HasPrefix(t, "regex") {
				res, err := processRegex(t, v.Field(i))
				if res == nil && err == nil {
					continue
				}
				return res, err
			}
		}
	}

	return nil, nil
}

func processRegex(tagStr string, field reflect.Value) (*ApiStatus, error) {
	targets := strings.SplitN(tagStr, "=", 2)
	if len(targets) != 2 {
		return nil, errors.New("tag 'min' error, miss argument")
	}
	if field.Kind() != reflect.String {
		return nil, fmt.Errorf("tag 'regex' unsupport type: %v", field.Kind())
	}
	ok, err := regexp.MatchString(targets[1], field.String())
	if err != nil {
		return nil, err
	}
	if !ok {
		return ErrErrorParam, nil
	}
	return nil, nil
}

func processTagMin(tagStr string, field reflect.Value) (*ApiStatus, error) {
	targets := strings.SplitN(tagStr, "=", 2)
	if len(targets) != 2 {
		return nil, errors.New("tag 'min' error, miss argument")
	}
	target := utils.StrToFloat(targets[1], -1)
	if target == -1 {
		return nil, fmt.Errorf("tag 'min' error, illegal argument %s", targets[1])
	}
	res, err := checkMin(field, target)
	if res == nil && err == nil {
		return nil, nil
	}
	return res, err
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
	default:
		return nil, fmt.Errorf("tag 'min' unsupport type: %v", val.Kind())
	}
	return nil, nil
}

func processTagMax(tagStr string, field reflect.Value) (*ApiStatus, error) {
	targets := strings.SplitN(tagStr, "=", 2)
	if len(targets) != 2 {
		return nil, errors.New("tag 'max' error, miss argument")
	}
	target := utils.StrToFloat(targets[1], -1)
	if target == -1 {
		return nil, fmt.Errorf("tag 'max' error, illegal argument %s", targets[1])
	}
	res, err := checkMax(field, target)
	if res == nil && err == nil {
		return nil, nil
	}
	return res, err
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
	default:
		return nil, fmt.Errorf("tag 'max' unsupport type: %v", val.Kind())
	}
	return nil, nil
}
