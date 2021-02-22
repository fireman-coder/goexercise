package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

//逆序+作为参数值传递（还有闭包引用哦）,4,3,2,1,0
