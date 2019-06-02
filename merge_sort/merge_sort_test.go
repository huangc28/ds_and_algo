package mergesort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	arr := []int{28, 99, 103, 94, 46}

	result := mergeSort(arr)

	assert.Equal(t, result, []int{28, 46, 94, 99, 103})
}

func TestAppendArray(t *testing.T) {
	arrA := []int{28, 99}
	arrA = arrA[1:]

	assert.Equal(t, arrA, []int{99})
}
