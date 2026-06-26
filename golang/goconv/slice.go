package goconv

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
