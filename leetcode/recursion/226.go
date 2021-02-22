package recursion

//226. 翻转二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	invert(root)
	return root
}

func invert(root *TreeNode) {
	if root == nil {
		return
	}
	//tmp := root.Left
	//root.Left = root.Right
	//root.Right = tmp
	root.Left, root.Right = root.Right, root.Left
	invert(root.Left)
	invert(root.Right)

}
