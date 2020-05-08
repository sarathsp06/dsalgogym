package arrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsOneAway(t *testing.T) {
	testCases := []struct {
		name   string
		input1 string
		input2 string
		output bool
	}{
		{"Equal Size", "", "", true},
		{"Equal Size", "a", "a", true},
		{"Equal Size", "a", "", true},
		{"Equal Size", "d", "a", true},
		{"Equal Size", "dbac", "dbgc", true},
		{"Equal Size", "dbac", "dbgx", false},
		{"Size Diff = 1", "a", "ab", true},
		{"Size Diff = 1", "cab", "ab", true},
		{"Size Diff = 1", "xyb", "ab", false},
		{"Size Diff = 1", "cab", "cd", false},
		{"Size Diff > 1", "", "ab", false},
		{"Size Diff > 1", "cab", "", false},
		{"Size Diff > 1", "cab", "cabby", false},
		{"CTCI", "pale", "ple", true},
		{"CTCI", "pale", "pales", true},
		{"CTCI", "pale", "bale", true},
		{"CTCI", "pale", "bake", false},
	}
	testFuncs := []func(string, string) bool{isOneAway}
	for _, fn := range testFuncs {
		for _, tc := range testCases {
			t.Run(fmt.Sprint(tc.name), func(t *testing.T) {
				assert.Equal(t, tc.output, fn(tc.input1, tc.input2))
			})
		}
	}
}
