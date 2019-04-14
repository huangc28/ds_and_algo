package insertionsort

// InsertionSort
func InsertionSortAscending(a []int) []int {
	for i := 2; i < len(a); i++ {
		nextNumToInsert := a[i]
		j := i - 1
		for j > 0 {
			if a[j] > nextNumToInsert {
				a[j+1], a[j] = a[j], nextNumToInsert
			}

			j--
		}
	}

	return a
}
