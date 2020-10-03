package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateMatrix(t *testing.T) {
	testcases := []struct {
		name   string
		input  [][]int
		output [][]int
	}{
		{"1x1 Matrix",
			[][]int{
				{0},
			},
			[][]int{
				{0},
			},
		},
		{"2x2 Matrix",
			[][]int{
				{1, 2},
				{3, 4},
			},
			[][]int{
				{2, 4},
				{1, 3},
			},
		},
		{"3x3 Matrix",
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[][]int{
				{3, 6, 9},
				{2, 5, 8},
				{1, 4, 7},
			},
		},
	}
	testFuncs := []func([][]int) [][]int{rotateMatrix}
	for _, fn := range testFuncs {
		for _, tc := range testcases {
			t.Run(tc.name, func(t *testing.T) {
				assert.Equal(t, tc.output, fn(tc.input))
			})
		}
	}
}
