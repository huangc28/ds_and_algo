package knapsack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKnapsanpRecursive(t *testing.T) {
	result := Knapsack(3, 50)

	assert.Equal(t, result, 220)
}

func TestKnapsanpBottomUp(t *testing.T) {
	result := KSBottomUp(3, 50)

	assert.Equal(t, result, 220)
}
