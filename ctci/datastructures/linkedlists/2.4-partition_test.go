package linkedlists

import (
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestPartition(t *testing.T) {
	testcases := []struct {
		name   string
		input  []int
		target int
	}{
		{"Empty", []int{}, 3},
		{"Single element", []int{200}, 3},
		{"Single element", []int{200}, 200},
		{"Two element", []int{200, 3}, 3},
		{"Two element", []int{3, 3}, 3},
		{"N element", []int{5, 4, 3, 2, 1}, 5},
		{"N element", []int{5, 4, 1, 5, 3, 7, 8}, 5},
		{"N element", []int{3, 5, 8, 5, 10, 2, 1}, 5},
	}
	testFuncs := []func([]int, int) string{partition}
	for _, fn := range testFuncs {
		for _, tc := range testcases {
			t.Run(tc.name, func(t *testing.T) {
				output := fn(tc.input, tc.target)
				outputEle := strings.Split(output, "-")
				seenTarget := false
				for _, val := range outputEle {
					ele, _ := strconv.Atoi(val)
					if seenTarget && ele < tc.target {
						t.Fatal("Failed - Partition unsuccessful ", outputEle)
					} else if ele == tc.target {
						seenTarget = true
					}
				}
				// Check that all the elements are still present
				ll := InitLinkedList(tc.input)
				inputEle := strings.Split(ll.String(), "-")
				sort.Strings(inputEle)
				sort.Strings(outputEle)
				if len(inputEle) != len(outputEle) {
					t.Fatal("Failed - Lengths do not match")
				}
				for i := range inputEle {
					if inputEle[i] != outputEle[i] {
						t.Fatal("Failed - Elements do not match")
					}
				}
			})
		}
	}
}
