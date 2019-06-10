package coinchange

func CoinChange(coins []int, amount int) int {
	var table [][]int

	if len(coins) == 0 || amount == 0 {
		return 0
	}

	// First construct a 2 dimensional array.
	// The x-axis to be the amount to construct
	// The y-axis to be the item / items to consider when constructing the table
	if table == nil {
		table = make([][]int, len(coins)+1)

		for i := range table {
			table[i] = make([]int, amount+1)
		}
	}

	maxVal := 0
	maxValKey := 0
	for i := 0; i <= len(coins); i++ {

		if i < len(coins) && coins[i] > maxVal {
			maxVal = coins[i]
			maxValKey = i
		}

		for a := 0; a <= amount; a++ {

			// If there is no item or no amount given
			// the number of combinatory coins is always 0
			if i == 0 || a == 0 {
				table[i][a] = -1
			} else {
				remainder := a % coins[i-1]

				if remainder == 0 {
					if a < coins[i-1] {
						table[i][a] = -1
					} else {
						// previous optimized solution
						table[i][a] = a / coins[i-1]
					}
				} else {
					pos := table[i-1][remainder]
					if pos == -1 {
						table[i][a] = -1
					} else {
						table[i][a] = (a / coins[i-1]) + pos
					}
				}
			}
		}
	}

	// Find the index of the largest coin
	return table[maxValKey+1][amount]
}
