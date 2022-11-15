package helper

func SumFloat64(arr []float64) float64 {
	var sum float64
	for _, val := range arr {
		sum += val
	}
	return sum
}

func SumFloat32(arr []float32) float32 {
	var sum float32
	for _, val := range arr {
		sum += val
	}
	return sum
}
