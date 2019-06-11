package maxprofit

func MaxProfit(prices []int) int {

	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if profit := prices[j] - prices[i]; profit > 0 {

			}
		}
	}

	return 0
}
