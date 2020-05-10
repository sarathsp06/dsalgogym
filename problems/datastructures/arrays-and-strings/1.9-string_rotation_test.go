package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsRotated(t *testing.T) {
	testcases := []struct {
		name   string
		input1 string
		input2 string
		output bool
	}{
		{"Empty String", "", "", true},
		{"Different Length", "", "a", false},
		{"Different Length", "ab", "a", false},
		{"Same Length ab", "ab", "ba", true},
		{"Same Length aa", "aa", "aa", true},
		{"Same Length waterbottle", "waterbottle", "bottlewater", true},
		{"Same Length waterbottle typo", "waterbottle", "bttlewatera", false},
		{"Same Length waterbottle typo", "abcd", "bcdd", false},
		{"Same Length waterbottle typo", "abcd", "bcda", true},
	}
	testFuncs := []func(string, string) bool{isRotated}
	for _, fn := range testFuncs {
		for _, tc := range testcases {
			t.Run(tc.name, func(t *testing.T) {
				assert.Equal(t, tc.output, fn(tc.input1, tc.input2))
			})
		}
	}
}
