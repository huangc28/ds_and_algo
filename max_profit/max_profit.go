package maxprofit

import (
	"math"
)

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

func Max(args ...int) int {
	maxVal := 0

	for _, v := range args {
		if v > maxVal {
			maxVal = v
		}
	}

	return maxVal
}

func Min(args ...int) int {
	minVal := math.MaxInt64

	for _, v := range args {
		if v < minVal {
			minVal = v
		}
	}

	return minVal
}

func MaxProfitDC(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	mid := len(prices) / 2
	rightHalf := prices[:mid]
	leftHalf := prices[mid:]

	case3 := Max(rightHalf...) - Min(leftHalf...)

	return Max([]int{MaxProfitDC(rightHalf), MaxProfitDC(leftHalf), case3}...)
}
