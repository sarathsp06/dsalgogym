package main

import (
	"fmt"
	"strconv"
)

// memoise this :D
func splitAsIP(s string, count int) []string {
	if !isValidIP(s, count) {
		return nil
	}

	if count == 1 {
		return []string{s}
	}

	count--
	splitAt := 3
	if s[0] == '0' {
		splitAt = 1
	} else if len(s) < splitAt {
		splitAt = len(s)
	}

	var splits []string
	for i := 1; i <= splitAt; i++ {
		if !isValidSplit(s[:i]) {
			continue
		}
		parts := splitAsIP(s[i:], count)
		for _, part := range parts {
			splits = append(splits, s[:i]+"."+part)
		}
	}
	return splits
}

func isValidSplit(s string) bool {
	i, err := strconv.Atoi(s)
	return err == nil && i < 256
}
func isValidIP(s string, count int) bool {
	return len(s) >= count && len(s) <= count*3
}

func main() {
	fmt.Println(splitAsIP("2542540123", 4))
}
