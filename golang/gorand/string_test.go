package gorand

import (
	"strings"
	"testing"
)

// TestGenerateRandomString_TableDriven проверяет длину и допустимые символы
func TestGenerateRandomString_TableDriven(t *testing.T) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Табличный тест (Table-Driven Test) - стандарт де-факто в Go
	tests := []struct {
		name   string
		length int
	}{
		{"Zero length", 0},
		{"Length 1", 1},
		{"Length 6 (default)", 6},
		{"Length 10", 10},
		{"Length 100", 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := GenerateRandomString(tt.length)

			// 1. Проверка длины
			if len(s) != tt.length {
				t.Errorf("Ожидалась длина %d, получено %d", tt.length, len(s))
			}

			// 2. Проверка набора символов
			for i, char := range s {
				if !strings.ContainsRune(charset, char) {
					t.Errorf("Символ %q на позиции %d отсутствует в допустимом наборе", char, i)
				}
			}
		})
	}
}

// TestGenerateRandomString_Unique проверяет, что генератор выдает разные строки
func TestGenerateRandomString_Unique(t *testing.T) {
	seen := make(map[string]bool)
	iterations := 10000 // 10 тысяч итераций

	for i := 0; i < iterations; i++ {
		s := GenerateRandomString(6)

		if seen[s] {
			// Вероятность коллизии для 6 символов (62^6 = ~56 млрд) при 10к итерациях ничтожна
			t.Fatalf("Обнаружен дубликат! Сгенерирована повторяющаяся строка: %s", s)
		}
		seen[s] = true
	}
}

// TestGenerateRandomString_NegativeLength проверяет поведение при отрицательной длине
// (Текущая реализация упадет с паникой, что нормально для Go, но мы это зафиксируем)
func TestGenerateRandomString_NegativeLength(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Ожидалась паника при отрицательной длине, но функция отработала без ошибок")
		}
	}()

	// Это должно вызвать панику (make([]byte, -1) или rand.IntN)
	_ = GenerateRandomString(-1)
}
