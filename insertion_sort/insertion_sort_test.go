package insertionsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSortDescendingOrder(t *testing.T) {
	speciments := []int{5, 104, 9, 7, 10, 29, 88, 22, 33, 46}

	sorted := InsertionSortAscending(speciments)

	assert.Equal(t, sorted, []int{5, 7, 9, 10, 22, 29, 33, 46, 88, 104})
}
