package goconv

import (
	"fmt"
	"regexp"
	"strconv"
)

func AnyToString(val any) (string, error) {
	if val == nil {
		return "", nil
	}
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

func ExtractNumberByIdPrefix(id, prefix string) (int, bool) {
	re := regexp.MustCompile(fmt.Sprintf(`^%s(\d+)$`, prefix))
	matches := re.FindStringSubmatch(id)
	if len(matches) != 2 {
		return 0, false
	}

	num, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		return 0, false
	}

	numInt, err := SafeInt64ToInt(num)
	if err != nil {
		return 0, false
	}
	return numInt, true
}
