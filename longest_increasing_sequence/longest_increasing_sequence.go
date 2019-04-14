package lis

import "log"

var table []int
var max int

// The Longest Increasing Subsequence (LIS) problem
// is to find the length of the longest subsequence
// of a given sequence such that all elements of the
// subsequence are sorted in increasing order.
// For example, the length of LIS for {10, 22, 9, 33, 21, 50, 41, 60, 80} is 6
// and LIS is {10, 22, 33, 50, 60, 80}.
// Dry Run, top bottom approach:
// 10 ---> 1
// 10;22 ---> 2
// 10;22
// 9 ---> 1
// 10;22;33 ---> 3
// 9;33 ---> 2
// 10;22;33 ---> 3
// 9:33
// table["10"] ---> 1
// table["10;22"] ---> 2
// func ArrayToStringKey(inputs []int) string {
// 	key := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(inputs)), ";"), "[]")

// 	return key
// }

func LIS(inputs []int) int {
	log.Printf("inputs %v", inputs)
	// retrieve the frist element of inputs.
	// convert the first element to key format
	// check if result table has the answer by using converted key.
	// compare the first element of inputs with n+1.
	// if table[n+1] > table[n], increase the result by 1 and store it in table
	// if table[n+1] <= table[n]
	if table == nil {
		table = make([]int, len(inputs))

		for i := range inputs {
			table[i] = 1
		}
	}

	for j := 0; j < len(inputs); j++ {
		i := 0
		for i < j {
			if inputs[i] < inputs[j] && table[i]+1 > table[j] {
				table[j] = table[i] + 1
			}

			if table[j] > max {
				max = table[j]
			}

			i++
		}
	}

	return max
}
