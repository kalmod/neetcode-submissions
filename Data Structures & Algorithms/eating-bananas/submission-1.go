
func minEatingSpeed(piles []int, h int) int {
	// this can be solved with bs
	// h = number of hours to eat naners. Our upper limit
	// ans = bananas to eat every hour <= h. but it's our upper bound
	// what is the l & r here? 1 and max pile size? 
	l := 1
	r := 0
	for _, pile := range piles {
		if r < pile {
			r = pile
		}
	}

	ans := r
	for l <= r {
		k := l + (r-l)/2
		if hoursToEatPiles(piles, k) <= h {
			ans = k
			r = k - 1
		} else {
			l = k + 1
		}
	}

	return ans
}

func hoursToEatPiles(piles []int, eatingSpeed int) int {
	h := 0
	for _, pile := range piles {
		h += int(math.Ceil(float64(pile) / float64(eatingSpeed)))
	}
	return h
}
