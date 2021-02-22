package recursion

import "fmt"

var result []string

//2n个格子，每个放(或)+剪枝
func generateParenthesis1(n int) []string {
	result = make([]string, 0)
	generate1(n, 0, 0, "")
	fmt.Println("result", len(result))
	return result
}

func generate1(n, left, right int, str string) {
	if (left+right)/2 == n {
		result = append(result, str)
		return
	}
	if left < n {
		generate1(n, left+1, right, str+"(")
	}
	if left > right {
		generate1(n, left, right+1, str+")")
	}
	return
}
