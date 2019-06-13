package maxsubarray

import (
	"log"
	"testing"
)

func TestMaxSubArrayCaseOne(t *testing.T) {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	// arr := []int{-2, 1}
	res := MaxSubArrayDP(arr)
	log.Printf("res %v", res)
}
