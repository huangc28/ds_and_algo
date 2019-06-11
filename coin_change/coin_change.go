package coinchange

func Min(num1, num2 int) int {
	if num1 <= num2 {
		return num1
	}

	return num2
}

func CoinChange(coins []int, amount int) int {
	numArr := make([]int, amount+1)
	numArr[0] = 0

	// Initialize the array to have
	for i := 1; i < len(numArr); i++ {
		numArr[i] = amount + 1
	}

	for a := 1; a <= amount; a++ {
		for c := 0; c < len(coins); c++ {
			if a >= coins[c] {
				tmp1 := numArr[a-coins[c]] + 1
				tmp2 := numArr[a]
				min := Min(tmp1, tmp2)

				numArr[a] = min
			}
		}
	}

	ways := numArr[amount]

	if ways > amount {
		ways = -1
	}

	return ways
}
