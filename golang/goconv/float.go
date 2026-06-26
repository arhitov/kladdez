package goconv

import (
	"fmt"
	"math"
)

func FloatToStr(f float64) string {
	if f == math.Trunc(f) {
		return fmt.Sprintf("%.0f", f) // Целое число
	}
	return fmt.Sprintf("%g", f) // Дробное число
}
