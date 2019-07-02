package lnrs

import "log"

func lengthOfLongestSubstringOne(s string) int {
	if len(s) == 0 {
		return 0
	}

	charMap := make(map[rune]int)
	curMax := 0
	srune := []rune(s)

	// We have 2 indexes, i and j. At first iteration,
	// "i" and "j" both  indicates the first element of the
	// string "s".
	i := 0

	// "j" would increases for each iteration.  Loop through strings.
	for j := 0; j < len(srune); j++ {
		// Check the map to see if the character s[j] has already exists in charMap

		// If s[j] does not exists in the charmap, store s[j] in the charmap.
		// The value of "i" remains the same
		if _, exists := charMap[srune[j]]; !exists {
			charMap[srune[j]] = j
		} else {
			for i <= j {
				pos, exists := charMap[srune[i]]

				if exists {
					delete(charMap, srune[i])

					if srune[i] == srune[j] {
						i = pos + 1
						break
					}

				}
				i++
			}

			charMap[srune[j]] = j
		}

		if len(charMap) > curMax {
			curMax = len(charMap)
		}
	}

	return curMax
}

func lengthOfLongestSubstringTwo(s string) int {
	if len(s) == 0 {
		return 0
	}

	srune := []rune(s)
	curMax, left := 0, 0

	for j := 0; j < len(s); j++ {
		if left < j && srune[left] == srune[j] {
			left = j + 1
		}

		curLen := (j - left) + 1

		if curLen > curMax {
			curMax = curLen
		}
	}

	log.Printf("cur max %v", curMax)

	return curMax
}
