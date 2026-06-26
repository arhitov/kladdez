package goconv

import "fmt"

const MMPerInch = 25.4 // 1 дюйм = 25.4 мм

// PxOnMm возвращает количество пикселей на 1 миллиметр при заданном DPI
func PxOnMm[T float64 | int](dpi T) float64 {
	if dpi <= 0 {
		panic(fmt.Errorf("DPI должен быть положительным числом"))
	}
	return float64(dpi) / MMPerInch
}

// MmToPx конвертирует миллиметры в пиксели
func MmToPx[T float64 | int](mm float64, dpi T) int {
	return int(mm * PxOnMm(dpi))
}

// MmToPt конвертирует миллиметры в пункты (PostScript/DTP points)
func MmToPt(mm float64) float64 {
	// 1 point = 1/72 дюйма → 1 мм = 72 / 25.4 points
	return mm * 72.0 / 25.4
}

// PxToMm конвертирует пиксели в миллиметры при заданном DPI
func PxToMm[T float64 | int, D float64 | int](px T, dpi D) float64 {
	if dpi <= 0 {
		panic(fmt.Errorf("DPI должен быть положительным числом"))
	}
	// мм = пиксели / (пикселей на мм) = пиксели / (DPI / 25.4) = пиксели × 25.4 / DPI
	return float64(px) * MMPerInch / float64(dpi)
}

// Вспомогательные функции для распространённых DPI

// PxToMmScreen — для стандартных экранов (96 DPI, Windows)
func PxToMmScreen[T float64 | int](px T) float64 {
	return PxToMm(px, 96)
}
