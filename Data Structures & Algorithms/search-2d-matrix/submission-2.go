func searchMatrix(matrix [][]int, target int) bool {
	// ascending order
	// matrix[i][len(matrix)-1] < matrix[i+1][0]
	// i think i can bs on row, then columns

	l := 0
	r := len(matrix)-1
	row_idx := 0
	for l <= r {
		var mid int = l + (r - l) / 2
		if target >= matrix[mid][0] {
			row_idx = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	l = 0
	r = len(matrix[row_idx]) - 1

	for l <= r {
		mid := l + (r - l) / 2
		if matrix[row_idx][mid] == target {
			return true
		} else if matrix[row_idx][mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return false
}
