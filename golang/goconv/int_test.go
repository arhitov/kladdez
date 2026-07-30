package goconv

import "testing"

func TestAnyToInt(t *testing.T) {
	tests := []struct {
		name         string
		input        any
		defaultValue int
		expected     int
	}{
		{"int", 42, 0, 42},
		{"float32", float32(3.14), 0, 3},
		{"float64", 2.718, 0, 2},
		{"string valid", "123", 0, 123},
		{"string invalid", "abc", 999, 999},
		{"bool true", true, 0, 1},
		{"bool false", false, 0, 0},
		{"*int valid", intPtr(100), 0, 100},
		{"*int nil", nil, 999, 999},
		{"nil value", nil, 999, 999},
		{"unsupported type", []int{1, 2, 3}, 999, 999},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AnyToIntOrDefault(tt.input, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("AnyToInt(%v, %d) = %d, expected %d",
					tt.input, tt.defaultValue, result, tt.expected)
			}
		})
	}
}

// Вспомогательная функция для создания указателя на int
func intPtr(i int) *int {
	return &i
}
