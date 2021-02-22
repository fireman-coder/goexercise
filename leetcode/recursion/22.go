package recursion

var res []string

func generateParenthesis(n int) []string {
	res = make([]string, 0)
	generate(0, 0, n, "")
	return res
}

//left括号<= n/2
//right<=left
func generate(left, right int, n int, str string) {
	if (left+right)/2 == n {
		res = append(res, str)
		return
	}
	//process1
	if left < n {
		//recursion
		generate(left+1, right, n, str+"(")
	}
	if right < left {
		generate(left, right+1, n, str+")")
	}

	//process2
	return
}
