package D_BFS

import "math"

var resMax []int
func largestValues(root *TreeNode) []int {
	resMax = make([]int,0)
	if root==nil{
		return  resMax
	}
	q:=[]*TreeNode{root}

	for len(q)!=0{
		p:=make([]*TreeNode,0)
		max:=math.MinInt64
		for _,v:=range q{
			if v.Val>max{
				max=v.Val
			}

			if v.Right!=nil{
				p=append(p,v.Right)
			}

			if v.Left!=nil{
				p=append(p,v.Left)
			}

		}

		resMax=append(resMax,max)
		q=p
	}

	return resMax
}

