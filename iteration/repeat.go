package iteration

func Repeat(x string, times int) string {
	result := ""
	for range times {
		result += x
	}
	return result
}
