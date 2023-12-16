package piscine

func RecursivePower(nb int, power int) int {
	if (nb < 0 && power == 0) || (nb > 0 && power == 0) {
		return 1
	}
	if (nb < 0 && power < 0) || (nb > 0 && power < 0) {
		return 0
	}
	if nb == 0 && power == 0 {
		return 1
	}
	result := nb
	return result * RecursivePower(nb, power-1)
}
