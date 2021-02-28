package main

import (
	"fmt"
	"unsafe"
)

func main(){
	a:=1

	//&a的地址值==unsafe.Pointer(&a)，两者都不能运算
	//unsafe.Pointer-->uintptr后可运算

	fmt.Println(&a)
	fmt.Printf("%v\n",unsafe.Pointer(&a))

	fmt.Println(uintptr(unsafe.Pointer(&a)))
	fmt.Println(uintptr(8))
}

