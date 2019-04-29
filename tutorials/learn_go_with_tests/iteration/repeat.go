package iteration

func Repeat(char string, x int) string {
	var output string
	for i := 0; i < x; i++ {
		output += char
	}
	return output
}
