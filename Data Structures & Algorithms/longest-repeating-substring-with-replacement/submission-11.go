func characterReplacement(s string, k int) int {
	// most frequent char in substring - length of substring

	m := make(map[rune]int)
	l, ans := 0, 0
	max_f := 0
	for idx, r := range s {
		m[r] += 1
		if m[r] > max_f {
			max_f = m[r]
		}
		for (idx-l+1)-max_f > k {
			m[rune(s[l])] -= 1
			l += 1
		} 

		if (idx-l+1) > ans {
			ans = (idx-l+1)
		}
		
	}
	return ans
}
