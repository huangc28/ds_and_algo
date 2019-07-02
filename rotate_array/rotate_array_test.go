package rotatearr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaseOne(t *testing.T) {
	testArr := []int{1, 2, 3, 4, 5, 6, 7}

	rotateImpv(testArr, 3)

	assert.Equal(t, []int{5, 6, 7, 1, 2, 3, 4}, testArr)
}
