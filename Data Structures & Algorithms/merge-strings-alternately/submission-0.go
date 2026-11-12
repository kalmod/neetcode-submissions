func mergeAlternately(word1 string, word2 string) string {
	array1 := []rune(word1)
	array2 := []rune(word2)
	l := 0
	r := 0

	combinedArr := []rune{}

	for l < len(array1)	&& r < len(array2) {
		combinedArr = append(combinedArr, array1[l])
		combinedArr = append(combinedArr, array2[r])
		l++
		r++
	}

	for l < len(array1) {
		combinedArr = append(combinedArr, array1[l])
		l++
	}

	for r < len(array2) {
		combinedArr = append(combinedArr, array2[r])
		r++
	}

	return string(combinedArr)
}
