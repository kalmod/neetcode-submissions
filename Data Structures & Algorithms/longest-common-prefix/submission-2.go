func longestCommonPrefix(strs []string) string {
    prefix := ""
	i := 0
	if len(strs[0]) == 0 {
		return prefix 
	}
	ch := string(strs[0][i])
	for {
		for _, str := range strs {
			if i >= len(str) || ch != string(str[i]) {
				return prefix
			}
		}
		prefix += ch
		i++
		if i >= len(strs[0]) {
			return prefix
		}
		ch = string(strs[0][i])
	}

	return prefix
}
