package rob

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	var table []int

	if table == nil {
		table = make([]int, len(nums)+1)
	}

	table[0] = 0
	table[1] = nums[0]
	maxRob := table[1]

	for i := 2; i <= len(nums); i++ {
		table[i] = max(table[i-2]+nums[i-1], table[i-1])

		if table[i] > maxRob {
			maxRob = table[i]
		}
	}

	return maxRob
}
