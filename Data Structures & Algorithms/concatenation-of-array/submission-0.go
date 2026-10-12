func getConcatenation(nums []int) []int {
	ans := make([]int, len(nums)*2)

	for i := 0; i < len(nums)*2; i++ {
		ans[i] = nums[i%len(nums)]
	}

	return ans
    
}
