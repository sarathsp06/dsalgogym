package watercapacity

import (
	"testing"
)

func Test_capacity(t *testing.T) {

	tests := []struct {
		arg  []int
		want int
	}{
		{[]int{5, 1, 2, 1, 3, 3, 4}, 10},
		{[]int{3, 7, 3, 1, 4, 8, 2}, 13},
		{[]int{9, 2, 5, 2, 1, 4, 8}, 26},
		{[]int{7, 4, 3, 3, 2, 1, 1}, 0},
		{[]int{7}, 0},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := capacity(tt.arg); got != tt.want {
				t.Errorf("capacity() = %v, want %v", got, tt.want)
			}
		})
	}
}
