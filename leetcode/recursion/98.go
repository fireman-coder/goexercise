package recursion

import "math"

//方法1：利用二叉搜索树的中序遍历，递增
var last *TreeNode

func isValidBST1(root *TreeNode) bool {
	last = &TreeNode{Val: -1 << 63}
	return inoder1(root)
}

func inoder1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !inoder1(root.Left) || root.Val <= last.Val {
		return false
	}
	last = root

	//if (root.Right != nil) && (!inoder1(root.Right) || root.Right.Val < root.Val) {
	//	return false
	//}
	//return true

	return inoder1(root.Right)
}

/*
最大值、最小值；

1、每个左边节点，
a、若其"祖宗节点"有左边节点则其小于其父节点，大于其“祖宗”节点
b、否则，小于其父节点，大于min

2、每个右节点，
a、若其"祖宗节点"有右边节点，大于其父节点，小于“祖宗节点”
b、否则，大于其父节点,小于max
*/

func isValidBST2(root *TreeNode) bool {
	return ValidBST2(root, math.MinInt64, math.MaxInt64)
}

func ValidBST2(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val > min && root.Val < max {
		return ValidBST2(root.Left, min, root.Val) &&
			ValidBST2(root.Right, root.Val, max)
	}
	return false
	//if root.Val <= min || root.Val >= max {
	//	return false
	//}
	//
	//if !ValidBST2(root.Left, min, root.Val) {
	//	return false
	//}
	//
	//if !ValidBST2(root.Right, root.Val, max) {
	//	return false
	//}
	//return true
}
