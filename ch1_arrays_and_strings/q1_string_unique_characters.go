package ch1_arrays_and_strings

// Page 90
//
// Implement an algorithm to determine
// if a string has all unique characters.

func IsStringUniqueChars(s string) bool {
	m := map[int32]struct{}{}
	for _, char := range s {
		_, ok := m[char]
		if ok {
			return false
		}
		m[char] = struct{}{}
	}
	return true
}
