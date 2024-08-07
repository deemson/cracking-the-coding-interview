package ch1_arrays_and_strings

import (
	"fmt"
	"unicode"
)

// Page 91
// Given a string, write a function to check if
// it is a permutation of a palindrome.

func IsPermutationOfPalindrome(s string) bool {
	characterCount := map[rune]int{}
	for _, r := range []rune(s) {
		if r == ' ' {
			continue
		}
		if !unicode.IsLetter(r) {
			panic(fmt.Sprintf("the character '%c' in a string '%s' is not a letter", r, s))
		}
		characterCount[r]++
	}
	isOddCountFound := false
	for _, count := range characterCount {
		if count%2 == 1 {
			if isOddCountFound {
				return false
			}
			isOddCountFound = true
		}
	}
	return true
}
