package recursion

/*
1、preorder 前序遍历，根左右，只要确定左、右子树的区间范围(preorder+inorder区间)，就可递归求出
2、在inorder中，根的左边是左子树的个数(preorder+个数得到preorder区间)+区间(inorder)，右边...
3、此题与297序列化与反序列化（解法：前序范遍历后-->递归即可）对比，因为297的前序结果若是叶子节点有2个nil，此条件就可分开左右子树（不需要分区为0作为递归的终止条件）
*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder)==0||len(preorder)==0{
		return nil
	}
	v:=preorder[0]
	root:=&TreeNode{Val: v}
	index:=findIndex(inorder,v)
	root.Left=buildTree(preorder[1:],inorder[:index])
	root.Right=buildTree(preorder[index+1:],inorder[index+1:])
	return root
}

func findIndex(inoder []int,target int)int{
	for i,v:=range inoder{
		if v==target{
			return i
		}
	}
	return -1
}