package fibonocci

// f0 = 0
// f1 = 1
// f2 = 1
// f3 = 2
// f4 = 3
// f5 = 5
// f6 = 8
func TopDownDPFib(num int) int {
	table := map[int]int{
		0: 0,
		1: 1,
	}

	if num <= 2 {
		return 1
	}

	if v, exists := table[num]; exists {
		table[num] = v
	} else {
		table[num] = TopDownDPFib(num-1) + TopDownDPFib(num-2)
	}

	return table[num]
}
