package maxprofit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfitDQ(t *testing.T) {
	res := MaxProfitDQ([]int{7, 1, 5, 3, 6, 4})

	assert.Equal(t, 5, res)
}

func TestMaxProfitCaseTwo(t *testing.T) {
	arr := []int{7, 6, 4, 3, 1}

	res := MaxProfit(arr)

	assert.Equal(t, 0, res)
}

func TestDCCaseOne(t *testing.T) {
	arr := []int{7, 1, 5, 3, 6, 4}

	res := MaxProfitDC(arr)

	assert.Equal(t, 6, res)
}
