package rotatearr

func rotate(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return nums
	}

	breakPoint := len(nums) - (k % len(nums))

	nums = append(nums[breakPoint:], nums[:breakPoint]...)

	return nums
}

func rotateImpv(nums []int, k int) {
	if len(nums) == 0 || k == 0 {
		return
	}

	n := len(nums)

	if k > len(nums) {
		k = k % len(nums)
	}

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start, end int) {
	i := start
	j := end

	for i < j {
		nums[i], nums[j] = nums[j], nums[i]

		i++
		j--
	}
}
