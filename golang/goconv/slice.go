package goconv

func AnyToSlice[O any](data any) []O {
	raw, ok := data.([]any)
	if !ok {
		return nil
	}
	result := make([]O, 0, len(raw))
	for _, valueRaw := range raw {
		if value, ok := valueRaw.(O); ok {
			result = append(result, value)
		}
	}
	return result
}

func ToInterfaceSlice[I any, O any](slice []I) []O {
	result := make([]O, len(slice))
	for i, v := range slice {
		result[i] = any(v).(O)
	}
	return result
}

func SliceAnyToSliceString[I interface{ String() string }](slice []I) []string {
	result := make([]string, len(slice))
	for i, v := range slice {
		result[i] = v.String()
	}
	return result
}
