func removeElement(nums []int, val int) int {
	k := 0
	for _, n := range nums {
		if n != val {
			nums[k] = n
			k+=1
		}
	}
	return k
}
