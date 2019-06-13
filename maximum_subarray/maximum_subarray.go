package maxsubarray

import (
	"log"
)

func sum(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func max(nums ...int) int {
	maxVal := 0

	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}

	return maxVal
}

func MaxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	mid := len(nums) / 2
	leftHalve := nums[:mid]
	rightHalve := nums[mid:]

	otherSol := sum(nums)

	return max([]int{MaxSubArray(leftHalve), MaxSubArray(rightHalve), otherSol}...)
}

func MaxSubArrayDP(nums []int) int {
	table := make([]int, len(nums))

	table[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		// We consider the following 3 circumstances for finding the maximum element
		//   1. i - 1
		//   2. i
		//   3. [i-1] + [i]
		table[i] = max([]int{nums[i], table[i-1] + nums[i]}...)
	}

	log.Printf("table %v", table)

	return 0
}
