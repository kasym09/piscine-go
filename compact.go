package piscine

func Compact(ptr *[]string) int {
	original := *ptr
	compactSize := 0

	// Count non-zero elements and compact the slice
	for _, value := range original {
		if value != "" {
			original[compactSize] = value
			compactSize++
		}
	}

	// Resize the slice to remove zero-valued elements
	*ptr = original[:compactSize]

	return compactSize
}
