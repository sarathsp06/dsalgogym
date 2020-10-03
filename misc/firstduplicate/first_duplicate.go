package firstduplicate

import (
	"math"
)

/**
Given an array of positive non-zero integers, such that maximum entry in the array is not
greater than the length of the array. Find the first duplicate entry in the array.

Minimize the time and the space complexity
*/

func firstDuplicate(input []int) int {
	var i, nextValue int
	for _, val := range input {
		i = int(math.Abs(float64(val))) - 1
		nextValue = input[i]
		if nextValue < 0 {
			return val
		}
		input[i] = -input[i]
	}
	return -1
}
