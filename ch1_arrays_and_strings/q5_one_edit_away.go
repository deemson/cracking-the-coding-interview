package ch1_arrays_and_strings

// Page 91
// There are three types of edits that can be performed on strings:
// insert a character, remove a character, or replace a character.
// Given two strings, write a function to check if they are
// one edit (or zero edits) away.

func StringsOneEditAway1(s1, s2 string) bool {
	switch {
	case len(s1) == len(s2):
		return stringsOneReplaceAway(s1, s2)
	case len(s1)+1 == len(s2):
		return stringsOneInsertAway(s1, s2)
	case len(s1)-1 == len(s2):
		return stringsOneInsertAway(s2, s1)
	}
	return false
}

func stringsOneReplaceAway(s1, s2 string) bool {
	runes1 := []rune(s1)
	runes2 := []rune(s2)
	foundDifference := false
	for index, rune1 := range runes1 {
		if rune1 != runes2[index] {
			if foundDifference {
				return false
			}
			foundDifference = true
		}
	}
	return true
}

func stringsOneInsertAway(s1, s2 string) bool {
	index1, index2 := 0, 0
	for index2 < len(s2) && index1 < len(s1) {
		if s1[index1] != s2[index2] {
			if index1 != index2 {
				return false
			}
			index2++
		} else {
			index1++
			index2++
		}
	}
	return true
}

func StringsOneEditAway2(s1, s2 string) bool {
	// make sure s1 is not longer than s2
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	index1 := 0
	index2 := 0
	foundDifference := false
	for index2 < len(s2) && index1 < len(s1) {
		if s1[index1] != s2[index2] {
			if foundDifference {
				return false
			}
			foundDifference = true
			if len(s1) == len(s2) {
				index1++
			}
		} else {
			index1++
		}
		index2++
	}
	return true
}
