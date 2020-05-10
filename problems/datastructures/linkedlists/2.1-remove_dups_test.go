package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsRotated(t *testing.T) {
	testcases := []struct {
		name   string
		input  []int
		output string
	}{
		{"Empty", []int{}, ""},
		{"Single Element", []int{1}, "1"},
		{"Single Element", []int{200}, "200"},
		{"No rep", []int{200, 3}, "200-3"},
		{"No rep", []int{200, 3, 4, 1, 2}, "200-3-4-1-2"},
		{"With rep ", []int{200, 3, 3, 1, 2}, "200-3-1-2"},
		{"With rep", []int{200, 3, 3, 1, 3}, "200-3-1"},
		{"With rep", []int{200, 3, 3, 3, 3}, "200-3"},
		{"With rep", []int{3, 3, 3, 3, 3}, "3"},
		{"With rep", []int{3, 1, 3, 1, 3}, "3-1"},
		{"With rep", []int{3, 1, 3, 1, 3, 2}, "3-1-2"},
	}
	testFuncs := []func([]int) string{removeDups, removeDupsNoBuffer}
	for _, fn := range testFuncs {
		for _, tc := range testcases {
			t.Run(tc.name, func(t *testing.T) {
				assert.Equal(t, tc.output, fn(tc.input))
			})
		}
	}
}
