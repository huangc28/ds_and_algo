package climbingstair

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbingRecurCaseOne(t *testing.T) {
	ways := ClimbingRecur(3)

	assert.Equal(t, 3, ways)
}

func TestClimbingDPCaseOne(t *testing.T) {
	ways := ClimbingDB(4)

	assert.Equal(t, 5, ways)
}
