package fibonocci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopDownDPFibonocci(t *testing.T) {
	res1 := TopDownDPFib(5)
	res2 := TopDownDPFib(6)
	res3 := TopDownDPFib(10)

	assert.Equal(t, res1, 5)
	assert.Equal(t, res2, 8)
	assert.Equal(t, res3, 55)
}
