package strutil

// Reverse takes a string and returns a reversed version
func Reverse(str string) (reversed string) {
	for _, c := range str {
		reversed = string(c) + reversed
	}
	return reversed
}
