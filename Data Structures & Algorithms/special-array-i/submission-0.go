func isArraySpecial(nums []int) bool {
	prevParity := nums[0] % 2
	for idx := 1; idx < len(nums); idx++ {
		if prevParity == nums[idx] % 2 {
			return false
		}
		prevParity = nums[idx] % 2
	}
	return true
}