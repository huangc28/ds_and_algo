package sortedarr2bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	// We first take the mid position.
	mid := len(nums) / 2

	tn := &TreeNode{}
	tn.Val = nums[mid]

	// We can be sure that every element on the right side of the mid position
	// are going to be larger than the element at mid since num is a sorted array.
	tn.Left = sortedArrayToBST(nums[:mid])
	tn.Right = sortedArrayToBST(nums[mid+1:])

	return tn
}
