package main

import "fmt"

/*
	接口的继承与转换
*/

//子集
type Humaner interface {
	SayHi()
}

//超集
type Personer interface {
	Humaner
	SayHello()
}

type Person struct {
	name string
	sex  string
}

type Student struct {
	Person
	score int
}

func (per *Person) SayHi() {
	fmt.Println("Hi")
}

func (stu *Student) SayHi() {
	fmt.Println("Hi~ (^_^)")
}

func (stu *Student) SayHello() {
	fmt.Println("Hello~ (^_^)")
}
func main() {

	//定义接口（子集）
	var h Humaner

	//实例化对象
	per := Person{"小明", "男"}
	h = &per

	//调用对象方法
	h.SayHi()

	//定义接口（超集）
	var p Personer

	stu := Student{Person{"小李", "男"}, 98}
	p = &stu

	//超集调用对象方法
	p.SayHi()
	p.SayHello()

	//子集调用对象方法
	h = &stu
	h.SayHi()

	//超集可以转换为子集，反之不可以
	h = p
	h.SayHi()

}
