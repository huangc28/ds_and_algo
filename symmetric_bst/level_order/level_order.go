package levelorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	table := make([][]int, 0, 1)
	var accuTable func(*TreeNode, int)

	accuTable = func(root *TreeNode, tl int) {
		if root == nil {
			return
		}

		if len(table) < tl {
			table = append(table, []int{})
		}

		// Add the root value to level order table
		table[tl-1] = append(table[tl-1], root.Val)

		// Increase the level indicator by 1
		accuTable(root.Left, tl+1)
		accuTable(root.Right, tl+1)
	}

	accuTable(root, len(table)+1)

	return table
}
