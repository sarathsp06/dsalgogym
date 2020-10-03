package minsteps

func minSteps(mat [][]bool, srcX, srcY, dstX, dstY int) (min int) {
	// dump algo - not A*
	switch {
	case srcX >= len(mat), srcY >= len(mat[0]),
		dstX >= len(mat), dstY >= len(mat[0]),
		srcX < 0, srcY < 0, dstX < 0, dstY < 0:
		return 0
	case mat[srcX][srcY]: // wall
		return 0

	}
	if (srcX == dstX) && (srcY == dstY) {
		min = 1
		return
	}
	// mark not already checked
	mat[srcX][srcY] = true

	// maybe do some A* thing here
	neighbours := [][]int{
		{srcX, srcY + 1},
		{srcX, srcY - 1},
		{srcX + 1, srcY},
		{srcX - 1, srcY},
	}
	min = 1000000000
	for _, node := range neighbours {
		minSteps := minSteps(mat, node[0], node[1], dstX, dstY)
		if minSteps > 0 && min > minSteps {
			min = minSteps
		}
	}
	if min != 1000000000 {
		mat[srcX][srcY] = false
		min = 1 + min
		return
	}
	return 0
}

// MinSteps finds minimum steps required to
// reach from source to destination in the given matrix mat
func MinSteps(mat [][]bool, srcX, srcY, dstX, dstY int) (min int) {
	res := minSteps(mat, srcX, srcY, dstX, dstY)
	if res == 0 {
		return 0
	}
	return res - 1
}
