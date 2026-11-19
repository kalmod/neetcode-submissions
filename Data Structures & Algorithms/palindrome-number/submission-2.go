func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	newNum := 0	
	tmp := x 
	for tmp != 0 {
		newNum = newNum * 10 + (tmp%10)
		tmp /= 10
	}
	// fmt.Println(x, tmp, newNum)
	return newNum == x
}
