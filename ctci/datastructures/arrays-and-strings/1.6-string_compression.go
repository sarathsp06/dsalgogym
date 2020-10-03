package arrays

import (
	"fmt"
	"strings"
)

// Complexity: O(n)
func compressString(input string) string {
	if len(input) == 0 {
		return input
	}

	src := []rune(input)
	prev := src[0]
	count := 0
	var output strings.Builder
	for _, ele := range []rune(input) {
		if ele != prev {
			fmt.Fprintf(&output, "%s%d", string(prev), count)
			prev = ele
			count = 1
		} else {
			count++
		}
	}
	fmt.Fprintf(&output, "%s%d", string(prev), count)

	fmt.Println(output.String())
	if len(output.String()) >= len(input) {
		return input
	}
	return output.String()
}
