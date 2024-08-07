package ch1_arrays_and_strings_test

import (
	"github.com/deemson/cracking-the-coding-interview/ch1_arrays_and_strings"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strconv"
	"strings"
	"testing"
)

func TestStringsOneEditAway(t *testing.T) {
	testCases := map[string]bool{
		"pale/ple":   true,
		"pales/pale": true,
		"pale/bale":  true,
		"pale/bae":   false,
	}
	for twoS, expected := range testCases {
		t.Run(twoS, func(t *testing.T) {
			s := strings.Split(twoS, "/")
			require.Len(t, s, 2)
			for index, f := range []func(string, string) bool{
				ch1_arrays_and_strings.StringsOneEditAway1,
				ch1_arrays_and_strings.StringsOneEditAway2,
			} {
				t.Run(strconv.Itoa(index+1), func(t *testing.T) {
					t.Run("forward", func(t *testing.T) {
						actual := f(s[0], s[1])
						assert.Equal(t, expected, actual)
					})
					t.Run("reversed", func(t *testing.T) {
						actual := f(s[1], s[0])
						assert.Equal(t, expected, actual)
					})
				})
			}
		})
	}
}
