package lis

import (
	"log"
	"testing"
)

func TestLongestIncreaseSequence(t *testing.T) {
	arr := []int{10, 22, 9, 33, 21, 50, 41, 60, 80}

	res := LIS(arr)

	log.Printf("res %v", res)
}
