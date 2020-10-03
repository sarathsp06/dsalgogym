package prime

import (
	"math"
)

// GetMaxPrimeFact computes and returns larget prime factor of the number
func GetMaxPrimeFact(n int64) int64 {
	x, ok := n, false
	limit := int64(math.Sqrt(float64(x)) + 1)
	s := make([]bool, limit)
	var i int64 = 2
	for ; i < limit; i++ {
		if s[i] {
			continue
		}
		for j := i; j < limit; j += i {
			s[j] = true
		}
		if x, ok = devide(x, i); ok {
			limit = int64(math.Sqrt(float64(x)) + 1)
		}

	}
	if x > i {
		return x
	}
	return i - 1
}

// devide devides largest integral multiple of y in x
// returns the devision result and if it was devided at least once info
func devide(x, y int64) (int64, bool) {
	var ok bool
	for x%y == 0 {
		x /= y
		ok = true
	}
	return x, ok
}
