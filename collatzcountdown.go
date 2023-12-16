package piscine

func CollatzCountdown(start int) int {
	counter := 0
	if start <= 0 {
		return -1
	}
	for i := start; i > 1; {
		counter++
		if i%2 == 0 {
			i = i / 2
		} else {
			i = i*3 + 1
		}
	}
	return counter
}
