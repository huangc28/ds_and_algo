package lps

import (
	"log"
	"testing"
)

func TestLPS(t *testing.T) {
	testStr := "abaabc"
	substring, num := LPA(testStr)

	log.Printf("substring %s, %d", substring, num)
}
