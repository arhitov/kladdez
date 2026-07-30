package gostring

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// PrefixBefore Возвращает строку до sep. Если sep не найден, возвращаёт строку полностью
func PrefixBefore(s, sep string) string {
	if i := strings.Index(s, sep); i != -1 {
		return s[:i]
	}
	return s
}

func FilterPrintableChars(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, str)
}

// TruncateString обрезает строку до заданной длины и добавляет "..."
// Если строка короче указанной длины, возвращает исходную строку
// Обрезает с начала строки, безопасно для Unicode
func TruncateString(s string, maxLength int) string {
	if maxLength <= 0 {
		return ""
	}

	// Если строка короче или равна максимальной длине, возвращаем как есть
	if utf8.RuneCountInString(s) <= maxLength {
		return s
	}

	// Преобразуем строку в срез рун для безопасной работы с Unicode
	runes := []rune(s)

	// Обрезаем до maxLength символов и добавляем "..."
	return string(runes[:maxLength]) + "..."
}

// TruncateBeginString обрезает строку с начала до заданной длины и ставит "..."
// Если строка короче указанной длины, оставляет как есть
// Безопасная обрезка Unicode
func TruncateBeginString(s string, maxLength int) string {
	if maxLength <= 0 {
		return "..."
	}

	// Если строка короче или равна максимальной длине, возвращаем как есть
	strLen := utf8.RuneCountInString(s)
	if strLen <= maxLength {
		return s
	}

	// Преобразуем строку в срез рун для безопасной работы с Unicode
	runes := []rune(s)

	// Берем последние maxLength символов и добавляем "..." в начало
	return "..." + string(runes[strLen-maxLength:])
}

// TruncateCenterString ограничивает длину строки, показывая начало и конец
// prefixLen - количество символов в начале
// suffixLen - количество символов в конце
func TruncateCenterString(s string, prefixLen, suffixLen int) string {
	totalLen := len(s)
	ellipsisLen := 3 // Длина многоточия "..."

	// Если строка короче или равна необходимой длине для отображения (учитывая многоточие)
	if totalLen <= prefixLen+suffixLen+ellipsisLen {
		return s
	}

	// Получаем начало и конец строки
	prefix := s[:prefixLen]
	suffix := s[totalLen-suffixLen:]

	return prefix + "..." + suffix
}

// TruncateCenterStringUnicode корректно работает с Unicode-символами
func TruncateCenterStringUnicode(s string, prefixLen, suffixLen int) string {
	runes := []rune(s)
	totalRunes := len(runes)
	ellipsisLen := 3 // Длина многоточия "..."

	// Если строка короче или равна необходимой длине для отображения (учитывая многоточие)
	if totalRunes <= prefixLen+suffixLen+ellipsisLen {
		return s
	}

	// Получаем начало и конец строки
	prefix := string(runes[:prefixLen])
	suffix := string(runes[totalRunes-suffixLen:])

	return prefix + "..." + suffix
}

// Substr
// NOTE: this isn't multi-Unicode-codepoint aware, like specifying skintone or
// gender of an emoji: https://unicode.org/emoji/charts/full-emoji-modifiers.html
func Substr(input string, start int, length int) string {
	asRunes := []rune(input)
	asRunesLen := len(asRunes)

	if start >= asRunesLen {
		return ""
	}
	// Если дилина меньше 0, считаем с конца
	if length < 0 {
		// Длина не может быть меньше 0
		length = max(asRunesLen+length-start, 0)
	}

	if start+length > asRunesLen {
		length = asRunesLen - start
	}

	return string(asRunes[start : start+length])
}
