package piscine

func UltimateDivMod(a *int, b *int) {
	// Convert to int64 for intermediate calculations
	numerator := int64(*a)
	denominator := int64(*b)

	// Store the result of the division in the int which a points to
	*a = int(numerator / denominator)

	// Store the remainder of the division in the int which b points to
	*b = int(numerator % denominator)
}
