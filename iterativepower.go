package piscine

func IterativePower(nb int, power int) int {
	result := nb
	if (nb < 0 && power == 0) || (nb > 0 && power == 0) {
		return 1
	}
	if (nb < 0 && power < 0) || (nb > 0 && power < 0) {
		return 0
	}
	if nb == 0 && power == 0 {
		return 1
	}
	for i := 1; i < power; i++ {
		result *= nb
	}
	return result
}
