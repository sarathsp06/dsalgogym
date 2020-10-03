package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompressString(t *testing.T) {
	testcases := []struct {
		name   string
		input  string
		output string
	}{
		{"Empty String", "", ""},
		{"Uncompressed", "a", "a"},
		{"Uncompressed", "aa", "aa"},
		{"Uncompressed", "aabb", "aabb"},
		{"Uncompressed", "aabbcc", "aabbcc"},
		{"Uncompressed", "abcabc", "abcabc"},
		{"Compressed aaa", "aaa", "a3"},
		{"Compressed aabcccccaaa", "aabcccccaaa", "a2b1c5a3"},
	}
	testFuncs := []func(string) string{compressString}
	for _, fn := range testFuncs {
		for _, tc := range testcases {
			t.Run(tc.name, func(t *testing.T) {
				assert.Equal(t, tc.output, fn(tc.input))
			})
		}
	}
}
