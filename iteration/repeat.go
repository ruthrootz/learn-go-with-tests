package iteration

func Repeat(x string) string {
	result := ""
	for range 5 {
		result += x
	}
	return result
}
