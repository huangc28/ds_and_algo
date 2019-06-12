package maxprofit

import "math"

func Max(values ...int) int {
	var maxVal int

	for _, elm := range values {
		if elm > maxVal {
			maxVal = elm
		}
	}

	return maxVal
}

func Min(values ...int) int {
	minVal := math.MaxInt64

	for _, elm := range values {
		if elm < minVal {
			minVal = elm
		}
	}

	return minVal
}

func MaxProfitDQ(prices []int) int {
	mid := len(prices) / 2

	if len(prices) <= 1 {
		return 0
	}

	rightHalve := prices[:mid]
	leftHalve := prices[mid:]

	// Find the max profit between rightHalve and leftHalve
	case3 := Max(leftHalve...) - Min(rightHalve...)

	return Max([]int{MaxProfitDQ(rightHalve), MaxProfitDQ(leftHalve), case3}...)
}
