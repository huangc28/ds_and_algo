package factorial

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBottomUpFactorial(t *testing.T) {
	res := SolveBottomUp(5)
	res2 := SolveBottomUp(6)

	assert.Equal(t, res, 120)
	assert.Equal(t, res2, 720)
}

func TestTopDownFactorial(t *testing.T) {
	res := SolveTopDown(5)
	res2 := SolveBottomUp(6)

	assert.Equal(t, res, 120)
	assert.Equal(t, res2, 720)
}
