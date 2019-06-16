package rob

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRobCaseOne(t *testing.T) {
	testArr := []int{1, 2, 3, 1}

	max := rob(testArr)
	assert.Equal(t, 4, max)
}

func TestRobCaseTwo(t *testing.T) {

	testArr := []int{2, 7, 9, 3, 1}

	max := rob(testArr)

	assert.Equal(t, 12, max)
}
func TestRobCaseThree(t *testing.T) {

	testArr := []int{}

	max := rob(testArr)

	log.Printf("max %v", max)
}

func TestRobCaseFour(t *testing.T) {

	testArr := []int{1, 2}

	max := rob(testArr)

	log.Printf("max %v", max)
}

func TestRobCaseFive(t *testing.T) {

	testArr := []int{2, 1, 1, 2}

	max := rob(testArr)

	log.Printf("max %v", max)
}
