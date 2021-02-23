package recursion

var result1 [][]int
func combine(n int, k int) [][]int {
	result1=make([][]int,0)
	generate3(1,n,k,make([]int,0))
	return result1
}

func generate3(index,n,k int,res []int){
	if k==0{
		res_bak:=make([]int,len(res))
		copy(res_bak,res)
		result1=append(result1,res_bak)
		return
	}

	for i:=index;i<=n;i++{
		if k>n-index+1{ //剪枝：如果剩下的n不够k选的了
			continue
		}

		res=append(res,i)
		generate3(i+1,n,k-1,res)
		res=res[:len(res)-1]
	}

}

