
type NumMatrix struct {
	m [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	n := NumMatrix{}
	n.m = make([][]int, len(matrix))	
	for i := range n.m {
		n.m[i] = make([]int, len(matrix[i])+1)
	}

	for i := 0; i < len(n.m); i++ {
			for j := 0; j < len(matrix[i]); j++ {
				n.m[i][j+1] = n.m[i][j] + matrix[i][j]
		}
	}

	
	
	return n
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	ans := 0
	for row := row1; row <= row2; row++ {
		ans = ans + (this.m[row][col2+1] - this.m[row][col1])
	}
	return  ans
	 
}

// Your NumMatrix object will be instantiated and called as such:
// obj := Constructor(matrix)
// param_1 := obj.SumRegion(row1,col1,row2,col2)
