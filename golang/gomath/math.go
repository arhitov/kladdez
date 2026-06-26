package gomath

func Min[T float32 | float64 | int](num ...T) T {
	var vMin T = num[0]
	for _, vNum := range num {
		if vMin > vNum {
			vMin = vNum
		}
	}
	return vMin
}
