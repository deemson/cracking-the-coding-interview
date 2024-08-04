package ch1_arrays_and_strings_test

import (
	"github.com/deemson/cracking-the-coding-interview/ch1_arrays_and_strings"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrlifyWhitespaces(t *testing.T) {
	testCases := map[string]string{
		"hello world":   "hello%20world",
		"aaa   bbb  cc": "aaa%20%20%20bbb%20%20cc",
	}
	for testCase, expected := range testCases {
		t.Run(testCase, func(t *testing.T) {
			actual := ch1_arrays_and_strings.UrlifyWhitespaces(testCase)
			assert.Equal(t, expected, actual)
		})
	}
}
