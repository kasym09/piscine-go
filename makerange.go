package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	result := createArrayofSize(min, max)
	return result
}

func createArrayofSize(min, max int) []int {
	size := max - min
	answer := make([]int, size)
	for i := 0; i < size; i++ {
		answer[i] = min + i
	}
	return answer
}
