func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	ans := nums[0]
	for l <= r {
		m := l + (r-l) / 2

		if nums[m] <= nums[l] && nums[m] <= nums[r] {
			if nums[m] < ans {
				ans = nums[m]
			}
			r = m - 1
		} else if nums[m] > nums[l] && nums[m] < nums[r] {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return ans
}
