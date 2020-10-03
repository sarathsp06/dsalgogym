package main

import (
	"fmt"
)

func tripletLinear(nums []uint) uint {
	var prev uint
	for i := 0; i < len(nums); i += 2 {
		if prev == nums[i] {
			break
		}
		prev = nums[i]
	}
	return prev
}

func triplet(nums []uint) uint {
	numLen := len(nums)
	// come up with a proper number after benchmark
	if numLen < 13 {
		return tripletLinear(nums)
	}

	mid := int(numLen / 2)
	if mid%2 == 0 {
		mid = mid - 1
	}
	if nums[mid] == nums[mid+1] {
		return triplet(nums[:mid+2])
	}
	return triplet(nums[mid+1:])
}

func main() {
	fmt.Println("Triplet: ", triplet([]uint{1, 1, 2, 2, 3, 3, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 34, 34, 35, 35, 45, 45, 43, 43, 89, 89, 89, 555, 555, 676, 676}))
}
