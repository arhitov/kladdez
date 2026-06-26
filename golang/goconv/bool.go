package goconv

import (
	"fmt"
)

func AnyToBoolWithOutError(v any) bool {
	r, _ := AnyToBool(v)
	return r
}

func AnyToBool(v any) (bool, error) {
	switch x := v.(type) {
	case bool:
		return x, nil
	case string:
		return x != "" && x != "0" && x != "false" && x != "False" && x != "FALSE", nil
	case int:
		return x != 0, nil
	case int8:
		return x != 0, nil
	case int16:
		return x != 0, nil
	case int32:
		return x != 0, nil
	case int64:
		return x != 0, nil
	case uint:
		return x != 0, nil
	case uint8:
		return x != 0, nil
	case uint16:
		return x != 0, nil
	case uint32:
		return x != 0, nil
	case uint64:
		return x != 0, nil
	case uintptr:
		return x != 0, nil
	case float32:
		return x != 0, nil
	case float64:
		return x != 0, nil
	case complex64:
		return x != 0, nil
	case complex128:
		return x != 0, nil
	case nil:
		return false, nil
	default:
		return false, fmt.Errorf("unsupported type [%T]", v)
	}
}
