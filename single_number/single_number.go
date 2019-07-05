package singlenum

import "log"

func singleNumber(nums []int) int {
	numMap := make(map[int]int)
	num := 0

	for i := 0; i < len(nums); i++ {
		_, exists := numMap[nums[i]]

		if !exists {
			numMap[nums[i]] = 1
		} else {
			numMap[nums[i]]++
		}
	}

	for k, v := range numMap {
		if v < 2 {
			num = k

			break
		}
	}

	return num
}

func singleNumberBitWise(nums []int) int {
	val := 0
	for i := 0; i < len(nums); i++ {
		val = val ^ nums[i]

		log.Printf("val at %d, %d", i, val)
	}

	return val
}
