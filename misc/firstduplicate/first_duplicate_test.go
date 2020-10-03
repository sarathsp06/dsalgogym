package firstduplicate
import (
	"testing"
)
func Test_firstDuplicate(t *testing.T){
	tests := []struct{
		args []int
		want int
	}{
		{[]int{1, 4, 2, 3}, -1},
		{[]int{2, 1, 2}, 2},
		{[]int{1, 3, 2, 4, 3, 2}, 3},
	}
	for _, tt := range tests{
		t.Run("testing first duplicate", func(t *testing.T){
			if got := firstDuplicate(tt.args); got != tt.want {
				t.Errorf("firstDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}