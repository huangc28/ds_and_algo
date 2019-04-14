package lcs

import (
	"log"
	"testing"
)

func TestLongestFindLCSInTopDownApproach(t *testing.T) {
	str1 := "ABCDGH"
	str2 := "AEDFHR"

	lcs := LCS(str1, str2, len(str1), len(str2))
	log.Printf("lcs %v", lcs)
}
