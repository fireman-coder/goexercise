package main

import (
"fmt"
"unsafe"
)

type Programmer struct {
	name string
	language string
}
//结构体会被分配一块连续的内存，结构体的地址也代表了第一个成员的地址。
func main() {
	p := Programmer{"stefno", "go"}
	fmt.Println(p)

	name := (*string)(unsafe.Pointer(&p))
	*name = "fireman"

	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.language)))
	*lang = "Golang"

	fmt.Println(p)
}
