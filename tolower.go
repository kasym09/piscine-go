package piscine

func ToLower(s string) string {
	bytelist := []byte(s)
	for i := 0; i < len(bytelist); i++ {
		if bytelist[i] >= 'A' && bytelist[i] <= 'Z' {
			bytelist[i] = bytelist[i] + 32
		}
	}
	return string(bytelist)
}
