package gostruct

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func StructToTagMap(s any, tagName string) (map[string]any, bool) {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	// Поддержка указателя
	if t.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil, false
		}
		v = v.Elem()
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil, false
	}

	result := make(map[string]any)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tagValue := field.Tag.Get(tagName)
		if tagValue == "" {
			continue // пропускаем поля без тега
		}
		result[tagValue] = v.Field(i).Interface()
	}

	return result, true
}

func ExtractTag(s any, fieldName string, tagName string) (string, bool) {
	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return "", false // не структура и не указатель на структуру
	}

	field, found := t.FieldByName(fieldName)
	if !found {
		return "", false
	}

	tag := field.Tag.Get(tagName)
	return tag, true
}

func ExtractTagData(s any, fieldName string, tagName string) (result map[string]string) {
	tag, ok := ExtractTag(s, fieldName, tagName)
	if !ok {
		return nil
	}
	parts := strings.Split(tag, ",")
	for _, part := range parts {
		expression := strings.Split(part, "=")
		if len(expression) == 2 {
			result[expression[0]] = expression[1]
		} else {
			result[expression[0]] = ""
		}
	}

	return
}

func FillingFields(
	dto any,
	tagName string,
	data map[string]string,
	customDecode func(valueStr, fieldType string) (any, error),
) error {
	// Используем рефлексию для автоматического присвоения значений
	dtoValue := reflect.ValueOf(dto).Elem()
	dtoType := dtoValue.Type()

	for i := 0; i < dtoValue.NumField(); i++ {
		field := dtoType.Field(i)
		tag := field.Tag.Get(tagName)
		if tag == "" {
			continue
		}

		// Парсим тег: "FieldName,Type"
		tagParts := strings.Split(tag, ",")
		if len(tagParts) != 2 {
			return fmt.Errorf("неверный формат тега для поля %s: %s", field.Name, tag)
		}

		fieldName := tagParts[0]
		fieldType := tagParts[1]

		valueStr, exists := data[fieldName]
		if !exists {
			continue
		}

		// Парсим значение в зависимости от типа
		var (
			parsedValue any
			err         error
		)
		if customDecode != nil {
			parsedValue, err = customDecode(valueStr, fieldType)
		}
		if err != nil {
			parsedValue, err = parseValueByType(valueStr, fieldType)
			if err != nil {
				return fmt.Errorf("ошибка парсинга поля %s (%s): %v", fieldName, fieldType, err)
			}
		}

		fieldValue := dtoValue.Field(i)
		parsedReflectValue := reflect.ValueOf(parsedValue)

		// Убеждаемся, что типы совместимы
		if parsedReflectValue.Type().AssignableTo(fieldValue.Type()) {
			fieldValue.Set(parsedReflectValue)
		} else {
			return fmt.Errorf("несовместимые типы для поля %s: ожидается %s, получено %s",
				fieldName, fieldValue.Type(), parsedReflectValue.Type())
		}
	}
	return nil
}

// parseValueByType парсит строку в значение определенного типа
func parseValueByType(valueStr, fieldType string) (any, error) {
	switch fieldType {
	case "string":
		return valueStr, nil

	case "bool":
		return parseBool(valueStr), nil

	case "time":
		return time.Parse(time.RFC3339, valueStr)

	case "int":
		return strconv.Atoi(valueStr)

	case "float":
		return strconv.ParseFloat(valueStr, 64)

	default:
		return nil, fmt.Errorf("неизвестный тип поля: %s", fieldType)
	}
}

// parseBool парсит строковое значение в bool
func parseBool(s string) bool {
	s = strings.ToLower(s)
	return s == "true" || s == "1" || s == "yes" || s == "y" || s == "on"
}
