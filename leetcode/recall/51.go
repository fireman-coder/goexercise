package recall

var queenRes [][]string
func solveNQueens(n int) [][]string {
	queenRes=make([][]string,0)
	res:=make([]string,0)

	queenBacktrack(n,0,res)
	return queenRes
}

func queenBacktrack(n,index int,res []string){
	if index==n{
		queenRes=append(queenRes,res)
		return
	}
	for  index<n{
		//res_bak:=make([]string,len(res))
		//copy(res_bak,res)
		index++
		res=append(res,position(n,index))
		queenBacktrack(n,index,res )
		res=res[:len(res)-1]
	}

}

func position(n,index int)string{
	str:=""
	for i:=0;i<n;i++{
		if i==index{
			str+="Q"
			continue
		}
		str+="."
	}
	return str
}
