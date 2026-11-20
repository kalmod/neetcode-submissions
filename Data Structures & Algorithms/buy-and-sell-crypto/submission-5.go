func maxProfit(prices []int) int {
	maxProfit := 0

	boughtPrice := prices[0]
	for _, price := range prices {
		if price < boughtPrice {
			boughtPrice = price
		} else {
			profit := price - boughtPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}
