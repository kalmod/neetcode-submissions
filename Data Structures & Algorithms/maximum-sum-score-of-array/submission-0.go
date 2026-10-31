func maximumSumScore(nums []int) int64 {
	postSum := make([]int, len(nums)+1)
	preSum := make([]int, len(nums)+1)

	for i := len(nums)-1; i >= 0; i-- {
		postSum[i] = postSum[i+1] + nums[i]
	}

	for i, n := range nums {
		preSum[i+1] = preSum[i] + n
	}

	ans := nums[0]

	for i := range nums {
		if ans < preSum[i+1] {
			ans = preSum[i+1]
		}
		if ans < postSum[i] {
			ans = postSum[i]
		}
	}

	return int64(ans)
}
