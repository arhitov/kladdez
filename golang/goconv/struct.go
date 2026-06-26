package goconv

import "encoding/json"

func AnyToStructViaJSON[T any](m any) (T, error) {
	var s T
	data, err := json.Marshal(m)
	if err != nil {
		return s, err
	}
	err = json.Unmarshal(data, &s)
	return s, err
}

func MapToStructViaJSON[T any](m map[string]any) (T, error) {
	var s T
	data, err := json.Marshal(m)
	if err != nil {
		return s, err
	}
	err = json.Unmarshal(data, &s)
	return s, err
}

// StructEnrichmentViaJSON Обогащает объект структуры данными из map через json теги
func StructEnrichmentViaJSON(s any, m map[string]any) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &s)
}
