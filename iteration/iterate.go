package iteration

// Repeat takes a char and repeats it n times
func Repeat(char string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated += char
	}
	return repeated
}