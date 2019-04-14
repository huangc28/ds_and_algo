package lcs

var table map[int]map[int]int

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// speciment: ABCDGH
// get all subsequences:
//  1. AB, AC, AD, AG, AH
//  2. BC,
func LCS(s1, s2 string, m, n int) int {

	result := 0
	rune1 := []rune(s1)
	rune2 := []rune(s2)

	if table == nil {
		table = make(map[int]map[int]int)
	}

	if _, exists := table[m]; !exists {
		table[m] = map[int]int{}
	}

	// m & n both are 0 in length
	// means we are dealing with
	// 2 empty strings.
	if m == 0 || n == 0 {
		return 0
	}

	if val, exists := table[m][n]; exists {
		return val
	}

	// the last character of s1 and s1
	// are the same. If they are the same,
	// we increase the the length by 1
	if rune1[m-1] == rune2[n-1] {
		result = 1 + LCS(s1, s2, m-1, n-1)
	}

	if rune1[m-1] != rune2[n-1] {
		subP1 := LCS(s1, s2, m-1, n)
		subP2 := LCS(s1, s2, m, n-1)

		result = Max(subP1, subP2)
	}

	table[m][n] = result

	return table[m][n]
}
