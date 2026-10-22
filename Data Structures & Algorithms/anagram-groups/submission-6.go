func groupAnagrams(strs []string) [][]string {
	h := make(map[[26]int][]string)

	for _, str := range strs {
		var freq [26]int
		for _, c := range str {
			freq[c-'a'] += 1
		}
		h[freq] = append(h[freq], str)
	}

	ans := [][]string{}
	for _, val := range h {
		ans = append(ans, val)
	}

	return ans
}
