/**
Given a 2d grid map of '1's (land) and '0's (water), count the number of islands.
An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
You may assume all four edges of the grid are all surrounded by water.
*/
package islandcount

import "testing"

func Test_islandcount(t *testing.T) {
	tests := []struct {
		args [][]int
		want int
	}{
		{[][]int{{1, 1, 1, 1, 0}, {1, 1, 0, 1, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 0, 0}}, 1},
		{[][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 1, 1}}, 3},
		{[][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, 0},
		{[][]int{{0, 1, 0}, {1, 0, 1}, {0, 1, 0}}, 4},
	}
	for _, tt := range tests {
		t.Run("testing numIslands", func(t *testing.T) {
			if got := numIslands(tt.args); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
