package lnrs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaseOne(t *testing.T) {
	str := "abcabcbb"

	len := lengthOfLongestSubstringOne(str)

	assert.Equal(t, 3, len)
}

func TestCaseTwo(t *testing.T) {
	str := "bbbbb"

	length := lengthOfLongestSubstringOne(str)

	assert.Equal(t, 1, length)
}

func TestCaseThree(t *testing.T) {
	str := "pwwkew"

	length := lengthOfLongestSubstringOne(str)

	assert.Equal(t, 3, length)
}

func TestCaseFour(t *testing.T) {
	str := "aab"

	length := lengthOfLongestSubstringOne(str)

	assert.Equal(t, 2, length)
}

func TestCaseFive(t *testing.T) {
	str := "qrsvbspk"

	length := lengthOfLongestSubstringOne(str)

	assert.Equal(t, 5, length)
}

func TestCaseSix(t *testing.T) {
	str := "dvdf"

	length := lengthOfLongestSubstringOne(str)

	assert.Equal(t, 3, length)
}
