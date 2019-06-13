package maxprofit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfitCaseOne(t *testing.T) {
	arr := []int{7, 1, 5, 3, 6, 4}

	res := MaxProfit(arr)

	assert.Equal(t, 5, res)
}

func TestMaxProfitCaseTwo(t *testing.T) {
	arr := []int{7, 6, 4, 3, 1}

	res := MaxProfit(arr)

	assert.Equal(t, 0, res)
}
