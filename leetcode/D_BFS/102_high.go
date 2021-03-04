package D_BFS

type TreeNode struct {
	    Val int
	    Left *TreeNode
	    Right *TreeNode
	 }


func levelOrder(root *TreeNode) [][]int {
	var res =make([][]int,0)
	if root==nil{
		return res
	}
	q:=[]*TreeNode{root}

	for len(q)!=0{ //for为了判断q!=nil
		r:=make([]int,0)
		p:=make([]*TreeNode,0)
		for j:=0;j<len(q);j++{ //for遍历这层的儿子节点
			r=append(r, q[j].Val)
			if q[j].Left!=nil{
				p=append(p,q[j].Left)
			}
			if q[j].Right!=nil{
				p=append(p,q[j].Right)
			}
		}
		q=p
		res=append(res,r)
	}
	return res
}


var res [][]int
func levelOrder_DFS(root *TreeNode) [][]int {
	res=make([][]int,0)
	if root==nil{
		return res
	}
	dfs(root,0)
	return res
}

func dfs(r *TreeNode,level int ){
	if r==nil{
		return
	}
	if level==len(res){
		res=append(res,[]int{})
	}
	res[level]=append(res[level],r.Val)
	dfs(r.Left,level+1)
	dfs(r.Right,level+1)
}
