package gomath

func Sum[T int | int8 | int16 | int32 | int64 | float32 | float64](nums []T) T {
	var sum T = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func Max[T float32 | float64 | int](num ...T) T {
	var vMax T = 0
	for _, vNum := range num {
		if vMax < vNum {
			vMax = vNum
		}
	}
	return vMax
}

func Min[T float32 | float64 | int](num ...T) T {
	var vMin T = num[0]
	for _, vNum := range num {
		if vMin > vNum {
			vMin = vNum
		}
	}
	return vMin
}
