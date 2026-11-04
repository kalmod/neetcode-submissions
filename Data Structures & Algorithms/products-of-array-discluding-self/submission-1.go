func productExceptSelf(nums []int) []int {
	pre := make([]int, len(nums)+1)
	post := make([]int, len(nums)+1)
	for i := range pre {
		pre[i] = 1
		post[i] = 1
	}

	for i := range nums {
		pre[i+1] = pre[i] * nums[i]
		post[len(nums)-1-i] = post[len(nums)-i] * nums[len(nums)-1-i]
	}

	res := []int{}

	for i := 0; i < len(pre)-1; i++ {
		res = append(res, pre[i]*post[i+1])
	}

	return res
}
