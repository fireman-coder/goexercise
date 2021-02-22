package recursion

//104二叉树最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return preorder(root, 0)
}

func preorder(root *TreeNode, count int) int {
	if root == nil {
		return 0
	}

	left := preorder(root.Left, count)
	right := preorder(root.Right, count)
	if left > right {
		return left + 1
	}
	return right + 1

}
