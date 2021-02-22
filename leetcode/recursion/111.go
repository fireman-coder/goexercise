package recursion

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)

	if (right == 0 || left < right) && left != 0 {
		return left + 1
	}
	return right + 1
}
