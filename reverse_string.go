package reverse_string

import "strings"

func ReverseString(input string) (output string) {
	if len(input) == 0 {
		return input
	}

	var sb strings.Builder
	var runes = []rune(input)

	for i := len(runes) - 1; i >= 0; i-- {
		sb.WriteRune(runes[i])
	}

	return sb.String()
}
