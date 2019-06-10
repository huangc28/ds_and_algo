package coinchange

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoinChangeCaseOne(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11

	num := CoinChange(coins, amount)

	assert.Equal(t, 3, num)
}

func TestCoinChangeCaseTwo(t *testing.T) {
	coins := []int{2}

	num := CoinChange(coins, 3)

	assert.Equal(t, -1, num)
}

func TestCoinChangeCaseThree(t *testing.T) {
	coins := []int{2, 5, 10, 1}
	amount := 27

	num := CoinChange(coins, amount)

	log.Printf("num %v", num)
}
