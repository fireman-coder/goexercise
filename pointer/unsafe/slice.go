package main

import (
	"fmt"
	"unsafe"
)
//https://qcrao91.gitbook.io/go/biao-zhun-ku/unsafe/ru-he-li-yong-unsafe-bao-xiu-gai-si-you-cheng-yuan
func main() {
	s := make([]int, 9, 20)
	var Len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(4)))
	fmt.Println(Len, len(s)) // 9 9

	var Cap = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	fmt.Println(Cap, cap(s)) // 20 20
}
