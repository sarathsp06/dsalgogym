package main

import (
	"fmt"
)

var romanValueMap = map[byte]int{
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	'V': 5,
	'I': 1,
}

// Assumption : the number is a valid roman number
func parseRoman(roman string) int {
	result, prevValue := 0, 0
	for i := len(roman) - 1; i >= 0; i-- {
		val := romanValueMap[roman[i]]
		if val < prevValue {
			result -= val
			continue
		}
		prevValue = val
		result += val
	}
	return result
}

func main() {
	var tests = map[string]int{
		"V":   5,
		"XXX": 30,
		"IX":  9,
		"XIV": 14,
	}

	for roman, expectedValue := range tests {
		fmt.Printf("Test(%v) %s: %d\n", parseRoman(roman) == expectedValue, roman, expectedValue)
	}
}
