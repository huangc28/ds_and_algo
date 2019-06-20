package merge2sortedarray

// merges nums1 into nums2 in the sorted desc manner.
func merge(nums1 []int, m int, nums2 []int, n int) {
	// We compare the last element of the two arrays to decide which to put at the given position
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]

			m--
		} else {
			nums1[m+n-1] = nums2[n-1]

			n--
		}
	}

	// Append the rest of the element from nums2 to nums1
	for n > 0 {
		nums1[m+n-1] = nums2[n-1]

		n--
	}
}
