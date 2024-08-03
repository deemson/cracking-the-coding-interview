package ch1_arrays_and_strings_test

import (
	"github.com/deemson/cracking-the-coding-interview/ch1_arrays_and_strings"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsStringUniqueChars(t *testing.T) {
	testCases := map[string]bool{
		"abc": true,
		"aaa": false,
	}
	for name, expected := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, expected, ch1_arrays_and_strings.IsStringUniqueChars(name))
		})
	}
}
