package solution

func maxProfit(prices []int, fee int) int {
	cash, hold := 0, -prices[0]
	for _, v := range prices[1:] {
		cash = max(cash, hold+v-fee)
		hold = max(hold, cash-v)
	}
	return cash
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
