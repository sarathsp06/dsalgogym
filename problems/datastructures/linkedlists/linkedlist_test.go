package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	testcases := []struct {
		name   string
		input  []int
		output string
	}{
		{"Empty", []int{}, ""},
		{"Single Element", []int{1}, "1"},
		{"Single Element", []int{200}, "200"},
		{"Multiple Elements", []int{200, 3}, "200-3"},
		{"Multiple Elements", []int{200, 3, 4, 1, 2}, "200-3-4-1-2"},
		{"Multiple Elements", []int{200, 3, 3, 1, 2}, "200-3-3-1-2"},
		{"Multiple Elements", []int{200, 3, 3, 1, 3}, "200-3-3-1-3"},
		{"Multiple Elements", []int{200, 3, 3, 3, 3}, "200-3-3-3-3"},
		{"Multiple Elements", []int{3, 3, 3, 3, 3}, "3-3-3-3-3"},
		{"Multiple Elements", []int{3, 1, 3, 1, 3}, "3-1-3-1-3"},
		{"Multiple Elements", []int{3, 1, 3, 1, 3, 2}, "3-1-3-1-3-2"},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			var ll LinkedList
			ll.Init(tc.input)
			assert.Equal(t, tc.output, ll.String())
		})
	}
}

func TestAdd(t *testing.T) {
	testcases := []struct {
		name   string
		input  []int
		elem   int
		output string
	}{
		{"Empty", []int{}, 3, "3"},
		{"Single Element", []int{1}, 2, "1-2"},
		{"Single Element", []int{200}, 3, "200-3"},
		{"Multiple Elements", []int{200, 3}, 4, "200-3-4"},
		{"Multiple Elements", []int{200, 3, 3, 3, 3}, 3, "200-3-3-3-3-3"},
		{"Multiple Elements", []int{3, 1, 3, 1, 3, 2}, 5, "3-1-3-1-3-2-5"},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			var ll LinkedList
			ll.Init(tc.input)
			ll.Add(tc.elem)
			assert.Equal(t, tc.output, ll.String())
		})
	}
}
