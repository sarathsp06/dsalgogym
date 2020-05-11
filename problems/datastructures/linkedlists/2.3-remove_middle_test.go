package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveMiddle(t *testing.T) {
	testcases := []struct {
		name   string
		input  []int
		target int
		output string
	}{
		{"With rep ", []int{200, 3, 3, 1, 2}, 3, "200-3-1-2"},
		{"With rep", []int{200, 3, 3, 1, 9}, 1, "200-3-3-9"},
		{"With rep", []int{1, 2, 3, 4, 5}, 2, "1-3-4-5"},
	}
	testFuncs := []func([]int, int) string{removeMiddle}
	for _, fn := range testFuncs {
		for _, tc := range testcases {
			t.Run(tc.name, func(t *testing.T) {
				assert.Equal(t, tc.output, fn(tc.input, tc.target))
			})
		}
	}
}
