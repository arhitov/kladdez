package primitive

import (
	"fmt"
	"github.com/arhitov/kladdez/golang/goconv"
)

// Size Размер
type Size struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

func NewSize(width, height float64) Size {
	return Size{Width: width, Height: height}
}

func (s Size) String() string {
	return fmt.Sprintf("%sx%s",
		goconv.FloatToStr(s.Width),
		goconv.FloatToStr(s.Height),
	)
}
