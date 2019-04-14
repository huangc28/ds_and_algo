package lis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestIncreaseSequence(t *testing.T) {
	arr := []int{10, 22, 9, 33, 21, 50, 41, 60, 80}

	res := LIS(arr)

	assert.Equal(t, res, 6)
}
