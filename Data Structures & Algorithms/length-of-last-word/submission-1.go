func lengthOfLastWord(s string) int {
	lastLength := 0
	idx := 0 
	for idx < len(s) {
		if s[idx] == ' ' {
			idx++
		} else {
			prevPos := idx
			
			for idx < len(s) && s[idx] != ' ' {
				idx++
			}
			lastLength = idx - prevPos
		}
	}
	return lastLength
}
