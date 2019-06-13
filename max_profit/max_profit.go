package maxprofit

import "math"

// Brute force
func MaxProfit(prices []int) int {
	var maxProfit int

	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if profit := prices[j] - prices[i]; profit > 0 && profit >= maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}

func Max(arr []int) int {
	val := 0

	for _, elem := range arr {
		if elem > val {
			val = elem
		}
	}

	return val
}

func Min(arr []int) int {
	val := math.MaxInt64

	for _, elem := range arr {
		if elem < val {
			val = elem
		}
	}

	return val
}

func MaxProfitDivideConquer(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	mid := len(prices) / 2
	rightHalf = prices[:mid]
	leftHalf = prices[mid:]

	case3 := Max(rightHalf) - Min(leftHalf)

	return 0
}
