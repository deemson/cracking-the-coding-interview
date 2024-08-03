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

func StringsArePermutation1(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	return sortString(s1) == sortString(s2)
}
