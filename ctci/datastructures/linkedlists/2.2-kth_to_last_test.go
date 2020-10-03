package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthToLast(t *testing.T) {
	testcases := []struct {
		name   string
		k      int
		input  []int
		output int
	}{
		{"Empty", 1, []int{}, -1},
		{"Single Element", 0, []int{1}, 1},
		{"Single Element", 1, []int{200}, -1},
		{"No rep", 1, []int{200, 3}, 200},
		{"No rep", 3, []int{200, 3, 4, 1, 2}, 3},
		{"With rep ", 5, []int{200, 3, 3, 1, 2}, -1},
		{"With rep ", 4, []int{200, 3, 3, 1, 2}, 200},
		{"With rep ", 3, []int{200, 3, 3, 1, 2}, 3},
		{"With rep ", 2, []int{200, 3, 3, 1, 2}, 3},
		{"With rep ", 1, []int{200, 3, 3, 1, 2}, 1},
		{"With rep ", 0, []int{200, 3, 3, 1, 2}, 2},
		{"With rep ", -1, []int{200, 3, 3, 1, 2}, -1},
	}
	testFuncs := []func(int, []int) int{kthToLast}
	for _, fn := range testFuncs {
		for _, tc := range testcases {
			t.Run(tc.name, func(t *testing.T) {
				assert.Equal(t, tc.output, fn(tc.k, tc.input))
			})
		}
	}
}
