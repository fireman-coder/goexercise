package recall

import "strings"

var queenRes [][]string
//var clos,na,pie map[int]int
/*
y=kx+b,k=+-1,y=rows,x=clos
clos
na=rows+clos
pie=rows-clos
*/

func solveNQueens(n int) [][]string {
	queenRes=make([][]string,0)
	pos:=make([][]string,n)
	clos,na,pie:=make(map[int]bool),make(map[int]bool),make(map[int]bool)
	for i,_:=range pos{
		line:=make([]string,0)
		for i:=0;i<n;i++{
			line=append(line,".")
		}
		pos[i]=line
	}
	queenBacktrack(n,0,pos,clos,na,pie)
	return queenRes
}

func queenBacktrack(n,rows int,pos [][]string,clos,na,pie map[int]bool){
	if n==rows{
		queenRes=append(queenRes,helper(pos))
		return
	}
		for clo:=0;clo<n;clo++{
			if  clos[clo]{
				continue
			}
			if  na[clo+rows]{
				continue
			}
			if  pie[rows-clo]{
				continue
			}
			clos[clo]=true
			na[clo+rows]=true
			pie[rows-clo]=true
			pos[rows][clo]="Q"
			queenBacktrack(n,rows+1,pos,clos,na,pie)
			pos[rows][clo]="."
			clos[clo]=false
			na[clo+rows]=false
			pie[rows-clo]=false
			//delete(clos,clo)
			//delete(na,clo+rows)
			//delete(pie,clo-rows)
		}
}

func helper(pos [][]string)[]string{
	var res []string
	for _,v:=range pos{
		res=append(res,strings.Join(v,""))
	}
	return res
}
