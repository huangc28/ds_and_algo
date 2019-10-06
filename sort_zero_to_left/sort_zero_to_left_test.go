package sortzerotoleft

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortZeroToLeft(t *testing.T) {
	nums := []int{0, 8, 9, 10, 0, 4, 9}

	res := sortZeroToLeft(nums)

	assert.Equal(t, res, []int{0, 0, 9, 10, 8, 4, 9})
}
