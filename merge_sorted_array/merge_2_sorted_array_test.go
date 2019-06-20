package merge2sortedarray

import (
	"log"
	"testing"
)

func TestCaseOne(t *testing.T) {
	t1 := []int{1, 2, 3, 0, 0, 0}
	t2 := []int{2, 5, 6}

	merge(t1, 3, t2, 3)
}

func TestCaseTwo(t *testing.T) {
	t1 := []int{4, 5, 6, 0, 0, 0}
	t2 := []int{1, 2, 3}

	merge(t1, 3, t2, 3)

	log.Printf("t1 %v", t1)
}
