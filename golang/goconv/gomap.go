package goconv

import (
	"fmt"
	"math"
	"strconv"
)

func ConvValue[T any](val any, def T) (T, error) {
	var zero T
	switch v := val.(type) {
	// Если тип соответствует
	case T:
		return v, nil
	// Преобразовываем из int в T
	case int:
		switch any(zero).(type) {
		case string:
			return any(strconv.Itoa(v)).(T), nil
		default:
			return val.(T), fmt.Errorf("unsupported type [%T] from int", zero)
		}
	// Преобразовываем из string в T
	case string:
		switch any(zero).(type) {
		case int:
			r, err := strconv.Atoi(v)
			if err != nil {
				return def, err
			} else {
				return any(r).(T), nil
			}
		case float64:
			r, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return def, err
			} else {
				return any(r).(T), nil
			}
		default:
			return val.(T), fmt.Errorf("unsupported type [%T] from string", zero)
		}
	// Преобразовываем из float64 в T
	case float64:
		switch any(zero).(type) {
		case int:
			return any(int(math.Round(v))).(T), nil
		default:
			return val.(T), fmt.Errorf("unsupported type [%T] from float64", zero)
		}
	default:
		return val.(T), fmt.Errorf("unsupported type [%T] to [%T]: %+v", val, zero, val)
	}
}

func ConvValueWithOutError[T any](val any, def T) T {
	v, _ := ConvValue[T](val, def)
	return v
}

// TruncateToDecimal обрезаем до указанной длины после запятой
func TruncateToDecimal(num float64, decimals int) float64 {
	factor := math.Pow(10, float64(decimals))
	return math.Round(num*factor) / factor
}
