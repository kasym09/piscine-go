package piscine

func ToUpper(s string) string {
	bytelist := []byte(s)
	for i := 0; i < len(bytelist); i++ {
		if bytelist[i] >= 'a' && bytelist[i] <= 'z' {
			bytelist[i] = bytelist[i] - 32
		}
	}
	return string(bytelist)
}
