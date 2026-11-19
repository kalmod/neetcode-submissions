func isPalindrome(x int) bool {
	newNum := 0	
	tmp := x 
	for tmp > 0 {
		newNum = newNum + (tmp%10)
		tmp /= 10
		if tmp == 0 {
			break
		}
		newNum *= 10
	}
	// fmt.Println(x, tmp, newNum)
	return x == newNum
}
