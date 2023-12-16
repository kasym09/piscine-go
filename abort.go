package piscine

func Abort(a, b, c, d, e int) int {
	list := []int{a, b, c, d, e}
	for i := 0; i < len(list); i++ {
		for j := 1; j < len(list); j++ {
			if list[j] > list[j-1] {
				list[j], list[j-1] = list[j-1], list[j]
			}
		}
	}
	return list[2]
}
