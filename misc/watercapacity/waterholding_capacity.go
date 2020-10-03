package watercapacity

import "fmt"

func capacity(elevations []int) int {
	if len(elevations) < 3 {
		return 0
	}
	var peaks []int
	for i, val := range elevations {
		top := len(peaks) - 1
		max := 0
		for j := top; j >= 0; j-- {
			if elevations[peaks[j]] >= val {
				top = j
				break
			}
			if max < elevations[peaks[j]] {
				top = j
				max = elevations[peaks[j]]
			}
		}
		peaks = append(peaks[:top+1], i)
	}
	fmt.Println(peaks)
	capacity := 0
	for i, j := 0, 1; j < len(peaks); i, j = i+1, j+1 {
		max := min(elevations[peaks[i]], elevations[peaks[j]])
		for p := peaks[i] + 1; p < peaks[j]; p++ {
			capacity = capacity + max - elevations[p]
		}
	}

	return capacity
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
