func characterReplacement(s string, k int) int {
	// most frequent char in substring - length of substring

	m := make(map[rune]int)
	l := 0
	frequent_ru := rune(s[0])
	ans := 0
	for idx, r := range s {
		m[r] += 1
		if m[r] > m[frequent_ru] {
			frequent_ru = r
		}
		if (idx-l+1)-m[frequent_ru] > k {
			m[rune(s[l])] -= 1
			l += 1
		} else {
			if (idx-l+1) > ans {
				ans = (idx-l+1)
			}
		}


	}

	return ans
}
