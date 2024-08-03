package ch1_arrays_and_strings_test

import (
	"github.com/deemson/cracking-the-coding-interview/ch1_arrays_and_strings"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestStringsArePermutation(t *testing.T) {
	testCases := map[string]bool{
		"listen/silent": true,
		"earth/heart":   true,
		"dog/cat":       false,
		"dog/book":      false,
	}
	for twoS, expected := range testCases {
		t.Run(twoS, func(t *testing.T) {
			s := strings.Split(twoS, "/")
			require.Len(t, s, 2)
			actual := ch1_arrays_and_strings.StringsArePermutation1(s[0], s[1])
			assert.Equal(t, expected, actual)
		})
	}
}
