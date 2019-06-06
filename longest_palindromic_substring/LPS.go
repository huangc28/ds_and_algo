package lps

// Declare table to be a two dimentional array.
var table [][]bool

func LPA(s string) (string, int) {
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
	// We need to declare a int variable for storing the largest number of substring.
	// We need declare a int variable for storing the longest  number of substring.

	largestNum := 0
	longestStringPalindrom := ""

	for l := 2; l < len(s); l++ {
		for i := 0; i <= len(s)-l; i++ {
			j := i + l - 1

			if srune[i] == srune[j] && table[i+1][j-1] == true {
				table[i][j] = true

				if j-i > largestNum {
					largestNum = j - i
					longestStringPalindrom = s[i : j+1]
				}
			}
		}
	}

	return longestStringPalindrom, len(longestStringPalindrom)
}
