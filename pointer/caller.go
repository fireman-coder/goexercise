package main

import "fmt"

/*
1、不管方法的接收者是什么类型，该类型的值和指针都可以调用，不必严格符合接收者的类型。
2、根据接收者的需求，编辑器来符合它
3、这个题没有接口的概念，只有实现接口A的所有方法的类型B，才能var A=B
4、reference：
https://qcrao91.gitbook.io/go/interface/zhi-jie-shou-zhe-he-zhi-zhen-jie-shou-zhe-de-qu-bie
*/
type Person struct {
	age int
}

func (p Person) howOld() int {
	return p.age
}

func (p *Person) growUp() {
	p.age += 1
}

func main() {
	// qcrao 是值类型
	qcrao := Person{age: 18}

	// 值类型, 接收者也是值类型的方法
	fmt.Println(qcrao.howOld()) //copy qcrao给接收者：(p Person) howOld()

	// 值类型, 接收者是指针类型的方法
	qcrao.growUp()    //&qcrao传给接收者 (p *Person) growUp()
	fmt.Println(qcrao.howOld()) //

	// ----------------------

	// stefno 是指针类型
	stefno := &Person{age: 100}

	// 指针类型, 接收者是值类型的方法
	fmt.Println(stefno.howOld())  //copy （*qcrao）给接收者：(p Person) howOld()

	// 指针类型, 接收者也是指针类型的方法
	stefno.growUp()  //qcrao指针传给接收者：(p Person) howOld()
	fmt.Println(stefno.howOld())
}
/*
18
19
100
101
*/