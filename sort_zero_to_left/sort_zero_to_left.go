package sortzerotoleft

func sortZeroToLeft(nums []int) []int {
	// if the length of nums slice equals 0, simply returns the first array
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}

	i := 0
	j := len(nums) - 1

	for i < j {
		if nums[j] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		} else {
			if nums[i] == 0 {
				i++
			} else {
				j--
			}
		}
	}

	return nums
}
