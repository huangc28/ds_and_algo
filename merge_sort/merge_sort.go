package mergesort

func merge(leftHalf []int, rightHalf []int) []int {
	mergedArr := make([]int, 0, len(leftHalf)+len(rightHalf))
	for len(leftHalf) > 0 || len(rightHalf) > 0 {
		if len(leftHalf) == 0 {
			return append(mergedArr, rightHalf...)
		}

		if len(rightHalf) == 0 {
			return append(mergedArr, leftHalf...)
		}

		if leftHalf[0] >= rightHalf[0] {
			mergedArr = append(mergedArr, rightHalf[0])
			// Push the first element out of the left half
			rightHalf = rightHalf[1:]
		} else {
			mergedArr = append(mergedArr, leftHalf[0])
			// Push the first element out of the right half
			leftHalf = leftHalf[1:]
		}
	}

	return mergedArr
}

func mergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	// Find the middle point to divide the array into two halves
	mid := len(slice) / 2

	// Call mergeSort for first half
	leftHalf := mergeSort(slice[:mid])

	// Call mergeSort for the second half
	rightHalf := mergeSort(slice[mid:])

	// Merge the two halves
	return merge(leftHalf, rightHalf)
}
