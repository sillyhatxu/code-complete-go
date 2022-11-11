package _21_Best_Time_to_Buy_and_Sell_Stock

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min, profit := prices[0], 0
	for _, price := range prices {
		if price < min {
			min = price
		} else if price-min > profit {
			profit = price - min
		}
	}
	return profit
}
