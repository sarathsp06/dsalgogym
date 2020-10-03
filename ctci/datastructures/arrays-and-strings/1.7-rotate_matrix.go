package arrays

// O(n^2) extra space / O(n^2) time
func rotateMatrix(input [][]int) [][]int {
	output := make([][]int, len(input))
	for i := range output {
		output[i] = make([]int, len(input))
	}
	for i := range input {
		for j := range input[i] {
			output[len(input)-1-j][i] = input[i][j]
		}
	}
	return output
}
