package intersectarray

import (
	"log"
)

func intersectMOne(nums1 []int, nums2 []int) []int {
	// Iterate through nums1. If element "x" at nums1 present in nums2,
	// place element "x" in the intersect array.
	ins := make([]int, 0, 1)
	charMap := make(map[int]int)

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			// If nums1[i] does presents in nums2, place nums1[i] in "ins" array
			if nums1[i] == nums2[j] {
				if pos, exists := charMap[nums1[i]]; exists {
					// nums1[i] has been previously found. Thus, we check the intersection at pos+1.
					if pos+1 < len(nums2) && nums1[i] == nums2[pos+1] {
						// If the nums2[pos+1] equals nums1[i], that means
						// j and j+1 are the same and appeared at consecutive position.
						charMap[nums1[i]] = pos + 1
						ins = append(ins, nums1[i])
					}
				} else {
					charMap[nums1[i]] = j
					ins = append(ins, nums1[i])
				}
				break
			}
		}
	}

	return ins
}

func intersectMTwo(nums1 []int, nums2 []int) []int {
	charMap := make(map[int]bool)
	intersect := make(map[int]bool)
	intersectArr := make([]int, 0, 1)

	for i := 0; i < len(nums1); i++ {
		charMap[nums1[i]] = true
	}

	for j := 0; j < len(nums2); j++ {
		if _, exists := charMap[nums2[j]]; exists {
			intersect[nums2[j]] = true
		}
	}

	log.Printf("intersect %v", intersect)

	for k := range intersect {
		intersectArr = append(intersectArr, k)
	}

	return intersectArr
}

func intersect(a1, a2 []int) []int {
	res := []int{}
	m1 := getInts(a1)
	m2 := getInts(a2)

	// m1 是较短的字典，会快一些
	if len(m1) > len(m2) {
		m1, m2 = m2, m1
	}
	// log.Printf("kkk %v", m1)
	for n := range m1 {
		// log.Printf("val %v %v %v %v", n, n, m1, m2)
		m1[n] = min(m1[n], m2[n])
	}

	// log.Printf("m1 %v", m1)
	// os.Exit(1)

	for n, size := range m1 {
		for i := 0; i < size; i++ {
			res = append(res, n)
		}
	}

	return res
}

// 清理重复的值，也便于查询
func getInts(a []int) map[int]int {
	res := make(map[int]int, len(a))

	for i := range a {
		res[a[i]]++
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
