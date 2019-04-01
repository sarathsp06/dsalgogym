package main

import (
	"errors"
	"fmt"
)

type SortedMatrix struct {
	matrix [][]int
	rCount int
	cCount int
}

// private so just trust :D
func (sm SortedMatrix) get(i, j int) int {
	return sm.matrix[i][j]
}

func NewSortedMatrix(mat [][]int) *SortedMatrix {
	sm := new(SortedMatrix)
	sm.matrix = mat
	sm.rCount = len(sm.matrix)
	if sm.rCount > 0 {
		sm.cCount = len(mat[0])
	}
	return sm
}

func (sm SortedMatrix) rowCount() int {
	return sm.rCount
}

func (sm SortedMatrix) columnCount() int {
	return sm.cCount
}

func (sm SortedMatrix) isValidIndex(i, j int) bool {
	if sm.rowCount() == 0 {
		return false
	}
	return i >= 0 && j >= 0 && i < sm.rowCount() && j < sm.columnCount()
}

func (sm SortedMatrix) CountSmaller(row, col int) (int, error) {
	if !sm.isValidIndex(row, col) {
		return 0, errors.New("invalid index")
	}
	val := sm.get(row, col)
	count := (row+1)*(col+1) - 1

	// row based
	for i, j := row-1, col+1; i >= 0 && j < sm.columnCount(); j++ {
		if sm.get(i, j) < val {
			count += i + 1
		} else {
			i--
			j--

		}
	}

	// column based
	for i, j := row+1, col-1; j >= 0 && i < sm.rowCount(); i++ {
		if sm.get(i, j) < val {
			count += j + 1
		} else {
			j--
			i--
		}
	}

	return count, nil
}

func (sm SortedMatrix) CountLarger(row, col int) (int, error) {
	if !sm.isValidIndex(row, col) {
		return 0, errors.New("invalid index")
	}
	val := sm.get(row, col)
	count := (sm.rowCount()-row)*(sm.columnCount()-col) - 1

	// row based
	for i, j := row+1, col-1; i < sm.rowCount() && j >= 0; j-- {
		if sm.get(i, j) > val {
			count += sm.rowCount() - i
			//fmt.Printf("sm[%d][%d] = %d, Count = %d\n", i, j, sm.get(i, j), count)
		} else {
			i++
			j++
		}
	}

	// column based
	for i, j := row-1, col+1; j < sm.columnCount() && i >= 0; i-- {
		if sm.get(i, j) > val {
			count += sm.columnCount() - j
			//fmt.Printf("sm[%d][%d] = %d, Count = %d\n", i, j, sm.get(i,j), count)
		} else {
			j++
			i++
		}
	}

	return count, nil
}

func main() {
	matrix := NewSortedMatrix([][]int{
		{1, 3, 7, 10, 15, 20},
		{2, 6, 9, 14, 22, 25},
		{3, 8, 10, 15, 25, 30},
		{10, 11, 12, 23, 30, 35},
		{20, 25, 30, 35, 40, 45},
	})

	smCount, _ := matrix.CountSmaller(1, 1)
	lrgCount, _ := matrix.CountLarger(3, 3)

	fmt.Printf("Count [(%d,%d) .. (%d,%d)] = %d ", 1, 1, 3, 3, smCount+lrgCount)
}
