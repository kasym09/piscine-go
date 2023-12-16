package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	result := createArrayofSize(min, max)
	return result
}
