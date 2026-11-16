func calPoints(operations []string) int { 
	stack := []int{}
	ans := 0
	for _, s := range operations {
		n, err := strconv.Atoi(s)
		if err != nil {
			if s == "+" {
				a := stack[len(stack)-1]
				b := stack[len(stack)-2]
				stack = append(stack, a+b)
				ans += (a+b)
			} else if s == "D" {
				ans += (stack[len(stack)-1] * 2)
				stack = append(stack, stack[len(stack)-1] * 2)
			} else if s == "C" {
				ans -= stack[len(stack)-1]
				stack = stack[:len(stack) - 1]
			}
		} else {
			stack = append(stack, n)
			ans += n
		}
	}


	return ans
}
