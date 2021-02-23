package recursion

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if  root==nil{
		return nil
	}
	if root==p||root==q{
		return root
	}
	left:=lowestCommonAncestor(root.Left,p,q)
	right:=lowestCommonAncestor(root.Right,p,q)
	//1、 p,q在root的左右分支上
	if left!=nil&&right!=nil{
		return root
	}
	/*
	2、p,q在“一条线上”
	3、p、q不存在
	*/
	if left==nil{
		return right
	}
	return left
}


