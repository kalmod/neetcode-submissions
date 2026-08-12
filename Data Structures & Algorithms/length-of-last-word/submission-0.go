func lengthOfLastWord(s string) int {
	lastLength := 0
	idx := 0 
	for idx < len(s) {
		prevPos := idx
		for idx < len(s) && s[idx] != ' ' {
			fmt.Printf("%c\n", s[idx])
			idx++
		}
		if idx != prevPos {
		lastLength = idx - prevPos

		}
		idx++	
	}
	return lastLength
}
