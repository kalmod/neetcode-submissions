func maxProfit(prices []int) int {
    profit := 0

    currMin := prices[0]

    for _, n := range prices {
        if n - currMin > profit {
            profit = n - currMin
        }

        if n < currMin {
            currMin = n
        }
    }
    return profit
}

