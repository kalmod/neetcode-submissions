func pivotIndex(nums []int) int {
	total := 0
	prefix := make([]int, len(nums)+1)
	for i, n := range nums {
		total += n
		prefix[i+1] = n + prefix[i]	
	}

	for i, n := range nums {
		rightSum := total - n - prefix[i]
		if prefix[i] == rightSum {
			return i
		}
	}

	return -1
	
}
