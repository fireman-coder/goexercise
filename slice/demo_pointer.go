
package main

import "fmt"

/*
type slice struct {
    array unsafe.Pointer // 元素指针
    len   int // 长度
    cap   int // 容量
}
pointer坑逼，指向地址可能不变，但地址内容会被修改，导致所有指向它的数据都被修改
*/
func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]  //2,3,4  长度为3，容量默认到数组末尾为8
	//s3:=s1[3]越界，s1[2:6:7]根据容量去判断的
	s2 := s1[2:6:7] //4,5,6,7 去索引2(闭)-6(开)；容量至索引7,为5

	s2 = append(s2, 100)  //4,5,6,7,100 来源slice
	//fmt.Println(&slice[4],&s2[0]) 地址相同
	s2 = append(s2, 200)  //4,5,6,7,100,200 创建新的slice

	s1[2] = 20

	fmt.Println(s1)  //2,3,20
	fmt.Println(s2) //4，5，6，7,100,200
	fmt.Println(slice)   //0, 1, 2, 3, 20, 5, 6, 7, 100, 9

}
