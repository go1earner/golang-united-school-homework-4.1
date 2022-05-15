package reverse_string

func ReverseString(input string) (output string) {
	runeStr := []rune(input)
	for i := len(runeStr); i != 0; i-- {
		output += string(runeStr[i-1])
	}
	return output
}
