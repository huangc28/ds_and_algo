package lps

import (
	"log"
	"testing"
)

func TestLPS(t *testing.T) {
	testStr := "abaabc"
	substring, num := LPS(testStr)

	log.Printf("substring %s, %d", substring, num)
}

func TestLPSCaseTwo(t *testing.T) {
	testStr := "abacab"
	substring, num := LPS(testStr)

	log.Printf("params %v, %v", substring, num)
}

func TestLPSCaseThree(t *testing.T) {
	testStr := "cbbd"

	substring, num := LPS(testStr)

	log.Printf("params %v, %v", substring, num)
}

func TestLPSCaseFour(t *testing.T) {
	testStr := "ccc"

	substring, num := LPS(testStr)

	log.Printf("params %v, %v", substring, num)
}
