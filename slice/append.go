package main

import "fmt"
/*
1、append注意点
2、s[len:]不算越界，s[len+1:]算越界
*/

func main(){
	nums := []int{1, 2, 3}

	fmt.Println(len(nums), cap(nums))
	for i, _ := range nums {
		a := append(nums[:i], nums[i+1:]...) //nums的 arrayPointer(指针不变，但指向内容会改变),len，cap 不变哦
		fmt.Println(cap(nums),len(nums))
		fmt.Println(nums, "----", i, "---", a)
	}

	nums = nums[3:] //不会算越界(index==len 不会算越界)
	//nums = nums[4 :] //会算越界


	//证明越界是通过len比较的
	n1:=make([]int,0,3)
	n1=append(n1,1)
	fmt.Println(n1[1:])   //index==len 不会算越界
	//fmt.Println(n1[2:]) //算越界 -->len
	//fmt.Println(nums[1]) //越界
}