package main

import "fmt"

func calc(index string,a,b int)int {
	ret:=a+b
	fmt.Println(index,a,b,ret)
	return ret
}
func main(){
	a:=1
	b:=2
	defer calc("1",a,calc("10",a,b))
	a=0
	defer calc("2",a,calc("20",a,b))
	b=1
}

/*
https://www.cnblogs.com/xinliangcoder/p/12079135.html
1、有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）
2、defer 调用的函数，参数的值在 defer 定义时就确定了:defer fmt.Println(a + b)
3、defer 函数内部所使用的变量的值需要在这个函数运行时才确定，看下代码：defer func() { fmt.Println(a + b) }()
4、当os.Exit()方法退出程序时，defer不会被执行。
5、defer 只对当前协程有效。
 */


/*正确答案
10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4

*/