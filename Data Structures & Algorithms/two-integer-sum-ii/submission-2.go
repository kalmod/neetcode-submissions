func twoSum(numbers []int, target int) []int {
	// sorted -> l is smallest, r is biggest num in array -> l, r = 0, len(arr)-1
	// if i want to make sum bigger, I move l pointer
	// if i want to make sum smaller, I move r pointer
	// answer always exists
	l := 0
	r := len(numbers)-1
	for l < r {
		if numbers[l] + numbers[r] < target {
			l++
		} else if numbers[l] + numbers[r] > target {
			r--
		} else {
			break
		}
	}
	return []int{l+1, r+1}
}
