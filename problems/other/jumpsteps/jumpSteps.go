package main

import (
	"fmt"
)

func getMaxValIdx(arr []int, first, last int) int {
	maxVal := arr[last]
	maxValIdx := last
	for idx, val := range arr[first+1 : last+1] {
		if val >= maxVal {
			maxVal = val
			maxValIdx = idx + first + 1
		}
	}
	return maxValIdx
}

func checkPath(arr []int) bool {
	for idx, val := range arr {
		arr[idx] = idx + val
	}
	idx := 0
	lastIdx := len(arr) - 1
	for {

		maxIdx := arr[idx]
		if maxIdx >= lastIdx {
			return true
		}
		if maxIdx == idx {
			return false
		}
		idx = getMaxValIdx(arr, idx, maxIdx)
	}
}

func main() {
	fmt.Println(checkPath([]int{1, 3, 1, 2, 0, 1}))
	fmt.Println(checkPath([]int{1, 2, 1, 0, 0, 0}))
	fmt.Println(checkPath([]int{0}))
}
