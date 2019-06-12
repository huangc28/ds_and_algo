package maxprofit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfitDQ(t *testing.T) {
	res := MaxProfitDQ([]int{7, 1, 5, 3, 6, 4})

	assert.Equal(t, 5, res)
}
