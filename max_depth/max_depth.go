package maxdepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var leftMaxDepth int
	var rightMaxDepth int

	if root.Left == nil && root.Right == nil {
		return 1
	} else if root.Left == nil && root.Right != nil {
		return 1 + maxDepth(root.Right)
	} else {
		return 1 + maxDepth(root.Left)
	}

	return max(leftMaxDepth, rightMaxDepth)
}
