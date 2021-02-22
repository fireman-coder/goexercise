package recursion

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums)==0{
		return nil
	}
	index,v:=max(nums)
	root:=&TreeNode{Val: v}
	root.Left=constructMaximumBinaryTree(nums[:index])

	root.Right=constructMaximumBinaryTree(nums[index+1:])
	return root
}

 func max(nums []int)(int,int){
 	if len(nums)==0{
 		return 0,0
	}
	index:=0
 	max:=nums[0]
 	for i,v:=range nums{
 		if v>max{
 			max=v
 			index=i
		}
	}
	return index,max
 }

