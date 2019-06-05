package lps

import "log"

// Declare table to be a two dimentional array.
var table [][]bool

func LPA(s string) {
	srune := []rune(s)

	if table == nil {
		table = make([][]bool, len(s))

		for i := range table {
			table[i] = make([]bool, len(s))
		}
	}

	// Initialize a two dimensional array.
	// A single character is always going to be a palindrome.
	// Thus, we set table[i][j] to be true where i === j.
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if i == j {
				table[i][j] = true
			} else {
				table[i][j] = false
			}
		}
	}

	// Determine whether substring with "length of 2" is palindromic or not
	for i := 0; i < len(s); i++ {
		j := i + 1

		if i < len(s)-1 {
			if srune[i] == srune[j] {
				table[i][j] = true
			} else {
				table[i][j] = false
			}
		}
	}

	// Determine whether substring with "length > 2" is palindromic or not
	// In this case "i" is the starting point of the substring
	for l := 2; l < len(s); l++ {
		for i := 0; i < len(s); i++ {
			j := i + l

			if srune[0] == srune[j] {

			}
		}
	}

	log.Printf("table %v", table)
}
