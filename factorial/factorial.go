package factorial

var (
	tableTopDown  map[int]int
	tableBottomUp map[int]int
)

func SolveTopDown(n int) int {
	if tableTopDown == nil {
		tableTopDown = make(map[int]int)
	}

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	tableTopDown[n] = n * SolveTopDown(n-1)

	return tableTopDown[n]
}

func SolveBottomUp(n int) int {
	if tableBottomUp == nil {
		tableBottomUp = make(map[int]int)
		tableBottomUp[0] = 0
		tableBottomUp[1] = 1
	}

	if v, exists := tableBottomUp[n]; exists {
		return v
	}

	for i := 2; i <= n; i++ {
		tableBottomUp[i] = tableBottomUp[i-1] * i
	}

	return tableBottomUp[n]
}
