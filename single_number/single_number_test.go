package singlenum

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaseOne(t *testing.T) {
	testArr := []int{4, 1, 2, 1, 2}

	num := singleNumber(testArr)

	assert.Equal(t, 4, num)
}

func TestCaseOneBitWise(t *testing.T) {

	testArr := []int{4, 1, 1}

	num := singleNumberBitWise(testArr)

	log.Printf("num %v", num)
	// assert.Equal(t, 4, num)
}
