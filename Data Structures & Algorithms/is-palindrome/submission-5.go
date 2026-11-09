func isPalindrome(s string) bool {
	chars := []byte(s)

	l := 0
	r := len(chars)-1
	for l < r {
		for l < r && isNotAlpha(chars[l]) {
			l++
		}

		for r > l && isNotAlpha(chars[r]) {
			r--
		}

		if l >= len(chars) || r < 0 || strings.ToLower(string(chars[l])) != strings.ToLower(string(chars[r])) {
			return false
		}
		l++
		r--
	}
	return true
}

func isNotAlpha(c byte) bool {
return !(c >= byte('a') && c <= byte('z')) && 
	!(c >= byte('A') && c <= byte('Z')) && 
	!(c >=  byte('0') && c  <= byte('9'))
}