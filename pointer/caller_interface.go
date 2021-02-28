package main

import "fmt"
/*
1、多了接口赋值
2、只有实现接口A的所有方法的类型B，才能var A=B，否则报错
3、方法接收者以值类型实现接口时，不管实体类型是值还是指针，都实现了该接口
方法接收者以指针实现接口时，只有实体类型是指针才能够实现该接口
interface是引用类型
4、reference
https://www.flysnow.org/2017/04/03/go-in-action-go-interface.html
*/

type coder interface {
	code()
	debug()
}

type Gopher struct {
	language string
}

func (p Gopher) code() {
	fmt.Printf("I am coding %s language\n", p.language)
}

func (p *Gopher) debug() {
	fmt.Printf("I am debuging %s language\n", p.language)
}

func main() {
	//Cannot use 'Gopher{"Go"}' (type Gopher) as type coder Type does not implement 'code'
	//as 'debug' method has a pointer receiver
	var c coder = Gopher{"Go"} //值类型Gopher，没有实现debug
	//解释：var c coder = Gopher是值运算，copy；(p *Gopher) debug()一般设计逻辑为改变其内容，如果语法可以，原值不会被改变。
	//var c coder = &Gopher{"Go"}  ok
	c.code()
	c.debug()
}
