package levelorder

import (
	"log"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	table := make([]int, 0, 1)

	table = append(table, 1, 2)

	log.Printf("table %v", table)
}
