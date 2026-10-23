func topKFrequent(nums []int, k int) []int {
	// first thing that comes to mind is a freq list
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}

	ans := make([][]int, len(nums))
	for n, v := range freq {
		ans[v-1] = append(ans[v-1], n)
	}

	res := []int{}
	i := len(ans)-1
	for i > -1 {
		if len(ans[i]) > 0 {
			for _, n := range ans[i] {
				res = append(res, n)
				if len(res) == k {
					return res
				}
			}	
		}
		i--
	}
	
	return res
}
