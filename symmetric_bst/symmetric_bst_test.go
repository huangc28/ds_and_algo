package symmetricbst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymmetricBSTIsValid(t *testing.T) {
	root := &TreeNode{
		Val: 0,
	}

	l21 := &TreeNode{
		Val: 2,
	}

	l22 := &TreeNode{
		Val: 2,
	}

	root.Left = l21
	root.Right = l22

	res := isSymmetric(root)

	assert.Equal(t, true, res)
}
