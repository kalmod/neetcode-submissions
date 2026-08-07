func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	ans := nums[0]
	for l <= r {
		m := l + (r-l) / 2

		if nums[m] < ans {
			ans = nums[m]
		}

		if nums[m] > nums[r] {
			// pivot is to the right
			l = m + 1
		} else {

			// pivot is to the left	
			r = m - 1
		}
	}

	return ans
}
