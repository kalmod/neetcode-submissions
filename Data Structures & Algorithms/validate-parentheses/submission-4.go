func isValid(s string) bool {
	r := []rune(s)
	res := []rune{}

	for _, ch := range r {
		n := len(res)
			if n > 0 && ch == '}' && res[n-1] == '{' {
			res = res[:n-1]
		} else if n > 0 && ch == ')' && res[n-1] == '(' {
			res = res[:n-1]
		} else if n > 0 && ch == ']' && res[n-1] == '[' {
			res = res[:n-1]
		} else {
			res = append(res, ch)
		}
	}

	if len(res) == 0 {
		return true
	} else {
		return false
	}
}
