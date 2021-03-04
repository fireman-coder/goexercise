package DP

func minDistance(word1 string, word2 string) int {
	n1,n2:=len(word1)+1,len(word2)+1

	dptable:=make([][]int,n1)

	for i:=range dptable{
		dptable[i]=make([]int,n2)
	}

	for i:=0;i<n1;i++{
		dptable[i][0]=i
	}

	for j:=0;j<n2;j++{
		dptable[0][j]=j
	}

	for i:=1;i<n1;i++{

		for j:=1;j<n2;j++{

			if word1[i-1]==word2[j-1]{
				dptable[i][j]=dptable[i-1][j-1]

			}else{
				dptable[i][j]=min(dptable[i-1][j],dptable[i][j-1],dptable[i-1][j-1])+1
			}



		}

	}


  return dptable[len(word1)][len(word2)]
}

func min(args... int)int{
	min:=args[0]
	for _,v:=range args{
		if v<min{
			min=v
		}
	}
	return min
}