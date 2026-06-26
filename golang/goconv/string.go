package goconv

import (
	"fmt"
	"strconv"
)

func AnyToString(val any) (string, error) {
	switch value := val.(type) {
	case float32:
		return fmt.Sprintf("%g", value), nil
	case float64:
		return fmt.Sprintf("%g", value), nil
	case int:
		return strconv.Itoa(value), nil
	case *int:
		return strconv.Itoa(*value), nil
	case string:
		return value, nil
	case *string:
		return *value, nil
	case bool:
		return strconv.FormatBool(value), nil
	default:
		return "", fmt.Errorf("неподдерживаемый тип: %T", val)
	}
}

func AnyToStringDefault(val any, def string) string {
	if v, err := AnyToString(val); err != nil {
		return def
	} else {
		return v
	}
}
