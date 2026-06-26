package goconv

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func AnyToIntOrDefault(v any, defaultValue int) int {
	if value, err := AnyToInt(v, defaultValue); err == nil {
		return value
	} else {
		return defaultValue
	}
}

func AnyToInt(v any, defaultValue int) (int, error) {
	if v == nil {
		return defaultValue, nil
	}

	switch value := v.(type) {
	case float32:
		return int(value), nil
	case float64:
		return int(value), nil
	case int:
		return value, nil
	case *int:
		if value == nil {
			return defaultValue, nil
		}
		return *value, nil
	case string:
		if value == "" {
			return defaultValue, nil
		} else if vInt, err := strconv.Atoi(value); err != nil {
			return defaultValue, fmt.Errorf("source [%T]%v: %v", v, v, err)
		} else {
			return vInt, nil
		}
	case bool:
		if value == true {
			return 1, nil
		} else {
			return 0, nil
		}
	default:
		return defaultValue, fmt.Errorf("неподдерживаемый тип: %T", v)
	}
}

func AnyToInt64(v any, defaultValue int64) (int64, error) {
	if v == nil {
		return defaultValue, nil
	}

	switch value := v.(type) {
	case float32:
		if value > math.MaxInt64 || value < math.MinInt64 {
			return defaultValue, fmt.Errorf("source [%T]%v: value out of int64 range", v, v)
		}
		return int64(value), nil
	case float64:
		if value > math.MaxInt64 || value < math.MinInt64 {
			return defaultValue, fmt.Errorf("source [%T]%v: value out of int64 range", v, v)
		}
		return int64(value), nil
	case int:
		return int64(value), nil
	case int8:
		return int64(value), nil
	case int16:
		return int64(value), nil
	case int32:
		return int64(value), nil
	case int64:
		return value, nil
	case uint:
		if value > math.MaxInt64 {
			return defaultValue, fmt.Errorf("source [%T]%v: value out of int64 range", v, v)
		}
		return int64(value), nil
	case uint8:
		return int64(value), nil
	case uint16:
		return int64(value), nil
	case uint32:
		return int64(value), nil
	case uint64:
		if value > math.MaxInt64 {
			return defaultValue, fmt.Errorf("source [%T]%v: value out of int64 range", v, v)
		}
		return int64(value), nil
	case *int64:
		if value == nil {
			return defaultValue, nil
		}
		return *value, nil
	case string:
		if value == "" {
			return defaultValue, nil
		}
		if vInt, err := strconv.ParseInt(value, 10, 64); err != nil {
			return defaultValue, fmt.Errorf("source [%T]%v: %w", v, v, err)
		} else {
			return vInt, nil
		}
	case bool:
		if value {
			return 1, nil
		}
		return 0, nil
	default:
		// Попытка обработать указатели на другие целочисленные типы
		val := reflect.ValueOf(v)
		if val.Kind() == reflect.Ptr && !val.IsNil() {
			return AnyToInt64(val.Elem().Interface(), defaultValue)
		}
		return defaultValue, fmt.Errorf("неподдерживаемый тип: %T", v)
	}
}

func SafeInt64ToInt(v int64) (int, error) {
	if v < math.MinInt || v > math.MaxInt {
		return 0, errors.New("value out of int range")
	}
	return int(v), nil
}
