func lengthOfLongestSubstring(s string) int {
	h := make(map[rune]bool)
	longestSubstringLength := 0
	start := 0

	for idx, r := range s {
		for h[r] {
			h[rune(s[start])] = false
			start++	
		}
		
		h[r] = true
		currLength := idx - start + 1
		if currLength > longestSubstringLength {
			longestSubstringLength = currLength
		}
	}
	return longestSubstringLength
	
}
