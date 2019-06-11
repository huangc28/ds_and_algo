package climbingstair

func ClimbingRecur(stairs int) int {
	ways := 0
	if stairs < 2 {
		ways = 1
	} else {
		ways = ClimbingRecur(stairs-1) + ClimbingRecur(stairs-2)
	}

	return ways
}

func ClimbingDB(stairs int) int {
	table := make([]int, stairs+1)
	table[0], table[1] = 1, 1

	for i := 2; i <= stairs; i++ {
		table[i] = table[i-1] + table[i-2]
	}

	return table[stairs]
}
