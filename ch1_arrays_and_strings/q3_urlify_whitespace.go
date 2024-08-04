package ch1_arrays_and_strings

// Page 90
// Write a method to replace all spaces
// in a string with '%20'.

func UrlifyWhitespaces(s string) string {
	whitespaceCount := 0
	for _, r := range s {
		if r == ' ' {
			whitespaceCount++
		}
	}
	runesIndex := len(s) + whitespaceCount*2
	runes := make([]rune, runesIndex)
	for stringIndex := len(s) - 1; stringIndex >= 0; stringIndex-- {
		if s[stringIndex] == ' ' {
			runes[runesIndex-1] = '0'
			runes[runesIndex-2] = '2'
			runes[runesIndex-3] = '%'
			runesIndex -= 3
		} else {
			runes[runesIndex-1] = rune(s[stringIndex])
			runesIndex--
		}
	}
	return string(runes)
}
