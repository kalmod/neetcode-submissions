func isPalindrome(s string) bool {
	chars := []byte(s)

	l := 0
	r := len(chars)-1
	for l < r {
		for l < r && !(chars[l] >= byte('a') && chars[l] <= byte('z')) && !(chars[l] >= byte('A') && chars[l] <= byte('Z')) && !(chars[l] >=  byte('0') && chars[l]  <= byte('9')) {
			l++
		}

		for r > l && !(chars[r] >= byte('a') && chars[r] <= byte('z')) && !(chars[r] >= byte('A') && chars[r] <= byte('Z')) && !(chars[r] >=  byte('0') && chars[r]  <= byte('9')) {
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
