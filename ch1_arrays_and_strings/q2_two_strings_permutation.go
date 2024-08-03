package ch1_arrays_and_strings

import "sort"

// Page 90
//
// Given two strings, write a method to decide if
// one is a permutation of the other.

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

// StringsArePermutation1 sorts letters in the strings and compares that the result is the same
func StringsArePermutation1(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	return sortString(s1) == sortString(s2)
}

// StringsArePermutation2 forms hash map of letter counts in s1
// and then counts down these numbers for each letter in s2.
// If one of the counts reaches negative value then there's difference.
func StringsArePermutation2(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	letters := make(map[rune]int, len(s1))
	for _, letter := range []rune(s1) {
		letters[letter]++
	}
	for _, letter := range []rune(s2) {
		letters[letter]--
		if letters[letter] < 0 {
			return false
		}
	}
	return true
}
