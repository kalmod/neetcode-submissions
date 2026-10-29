func maxScore(s string) int {
    rightSum := make([]int, len(s)+1)
    leftSum := make([]int, len(s)+1)
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if runes[len(runes)-i-1] == '1' {
			leftSum[len(runes)-i-1] = leftSum[len(runes)-i] + 1
		} else {
			leftSum[len(runes)-i-1] = leftSum[len(runes)-i] 
		}
		if runes[i] == '0' {
			rightSum[i+1] = rightSum[i] + 1
		} else {
			rightSum[i+1] = rightSum[i]
		}
		
	}

	res := 0

	for i := 1; i < len(s); i++ {
		if res < leftSum[i] + rightSum[i] {
			res = leftSum[i] + rightSum[i]
		}
	}


	return res
}