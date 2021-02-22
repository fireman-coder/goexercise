package main

import (
"io"
"log"
)




func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)  // 7, EOF
	}()
	return 7, io.EOF // 7, EOF
}

func func2(s string) ( int,  error) {
	n:=7
	defer func() {
		n+=1
		log.Printf("func1(%q) = %d, %v", s, n, nil)  //8 nil
	}()
	return n, io.EOF  // 7, EOF
}

//先return后defer
//原因就是return会将返回值先保存起来，对于无名返回值来说，
//保存在一个临时对象中，defer是看不到这个临时对象的；
//而对于有名返回值来说，就保存在已命名的变量中。
//https://juejin.im/post/6844903877033066509

func func3(s string) (n int, err error) {

	defer func() {
		n+=1
		log.Printf("func1(%q) = %d, %v", s, n, err) //8, EOF
	}()
	return 7, io.EOF //8, EOF
}

func main() {
	n,err:=func3("Go")
	log.Printf("----%d, %v",  n, err)

}