package knapsack

var valueBU = []int{
	60,
	100,
	120,
}

// The first element "0" is just the dummy element.
var weightBU = []int{
	10,
	20,
	30,
}

// The bottom up dynamic programming approach
func KSBottomUp(index, capacity int) int {
	var table [][]int

	//
	if table == nil {
		table = make([][]int, index+1)

		for i := range table {
			table[i] = make([]int, capacity+1)
		}
	}

	// "i" indicates the item of value and weight pair. For example:
	//   - 0 indicates that we are not considering any item.
	//   - 1 indicates that we are considering 1 item. In this case (weight, value)(10, 60)
	//   - 2 indicates that we are considering 2 items (10, 60) and (20, 100)
	for i := 0; i <= index; i++ {
		for w := 0; w <= capacity; w++ {
			if i == 0 || w == 0 {
				table[i][w] = 0
			} else if w >= weightBU[i-1] {
				// if the current capacity "w" can accomodate the current iterated weight,
				// then we will have two options we can decided upon:
				//   - Accomodate the current item into the knapsack. In this case we have to deduct
				//     The current weight from the weight of the item.
				//   - Do not accomodate the current item into the knapsack

				table[i][w] = Max(valueBU[i-1]+table[i-1][w-weightBU[i-1]], table[i-1][w])
			} else {
				table[i][w] = table[i][w-1]
			}
		}
	}

	return table[index][capacity]
}
