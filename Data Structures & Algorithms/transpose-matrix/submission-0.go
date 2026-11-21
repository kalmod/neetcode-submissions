func transpose(matrix [][]int) [][]int {
	row := len(matrix)	
	col := len(matrix[0])
	newMatrix := make([][]int, col)
	for idx := range newMatrix {
		newMatrix[idx] = make([]int, row)
	}

	for r := range matrix {
		for c := range matrix[r] {
			newMatrix[c][r] = matrix[r][c]
		}
	}
	return newMatrix
}
