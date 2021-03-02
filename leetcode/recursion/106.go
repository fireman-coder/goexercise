package recursion

//106. 从中序与后序遍历序列构造二叉树
/*
递归：
1、根据后序可确定根节点
2、根据中序可确定左/右子树区间（左子树个数）
3、父问题-->子问题
*/
func buildTree_106(inorder []int, postorder []int) *TreeNode {
	if len(postorder)==0{
		return nil
	}
	v:=postorder[len(postorder)-1]
	index:=findIndex_106(inorder,v)
return &TreeNode{
	Val: v,
	Left:buildTree_106(inorder[:index],postorder[:index]) ,
	Right:buildTree_106(inorder[index+1:],postorder[index:len(postorder)-1]) ,
}
}
//在inorder找到的根节点index=左子树的个数
func findIndex_106(inorder []int,target int )int{
	for i,v:=range inorder{
		if v==target{
			return i
		}
	}
	return -1
}