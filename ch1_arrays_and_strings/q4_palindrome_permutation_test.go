package ch1_arrays_and_strings_test

import (
	"github.com/deemson/cracking-the-coding-interview/ch1_arrays_and_strings"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPermutationOfPalindrome(t *testing.T) {
	testCases := map[string]bool{
		"never odd or even": true,
		"asd":               false,
	}
	for testCase, expected := range testCases {
		t.Run(testCase, func(t *testing.T) {
			actual := ch1_arrays_and_strings.IsPermutationOfPalindrome(testCase)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestIsPermutationOfPalindrome_Panic(t *testing.T) {
	defer func() {
		msg := recover()
		assert.Equal(t, "the character '#' in a string 'a#' is not a letter", msg)
	}()
	ch1_arrays_and_strings.IsPermutationOfPalindrome("a#")
}
