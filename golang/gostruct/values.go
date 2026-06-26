package gostruct

import "reflect"

func ExtractValue(s any, fieldName string) (any, bool) {
	v := reflect.ValueOf(s)

	// Разыменовываем указатель, если нужно
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil, false
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, false
	}

	field := v.FieldByName(fieldName)
	if !field.IsValid() {
		return nil, false
	}

	return field.Interface(), true
}
