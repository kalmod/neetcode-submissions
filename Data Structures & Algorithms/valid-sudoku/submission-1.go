func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, len(board))
	cols := make([]map[byte]bool, len(board[0]))
	box := make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		box[i] = make(map[byte]bool)
	}

	for i, row := range board {
		for j, val := range row {
			if val == '.' {
				continue
			}
			if _, ok := cols[j][val]; ok {
				// col value exists in column
				return false
			} else {
				cols[j][val] = true	
			}
			if _, ok := rows[i][val]; ok {
				// col value exists in row
				return false
			} else {
				rows[i][val] = true
			}
				
			boxIdx := (i/3)*3 + j/3
			if _, ok := box[boxIdx][val]; ok {
				return false
			} else {
				box[boxIdx][val] = true
			}

		}
	}


	return true

}