package main

import "fmt"

/*
	方法值和方法表达式（方法调用方式）
*/

type Person struct {
	name string
	sex  string
	addr string
}

type Student struct {
	Person
	score int
}

//指针类型作为接收者
func (s *Student) test() {
	fmt.Println("My test")
}

//将函数类型作为函数参数
func demo(f func()) {
	fmt.Println("哈哈~")

}
func main() {
	var stu Student
	//stu := Student{Person{"小明", "男", "北京"}, 98}

	//对象方法在代码区的存储位置，类型为函数
	fmt.Printf("%p\n", stu.test)
	fmt.Printf("%T\n", stu.test)

	//传统调用方式：
	stu.test()

	//定义函数类型的变量
	var f func()

	//调用方式2：方法值(方法地址)
	f = stu.test //将对象的方法地址赋值给函数类型变量
	f()          //已隐藏接收者，调用时无需再传递接收者，等价于stu.test()

	//调用方式3：方法表达式（）
	f2 := (*Student).test //将指定类型的函数赋值给变量
	f2(&stu)              //采用显性方式将接收者传递给函数

	//将函数类型作为参数传递
	demo(stu.test)
	demo(f)
}
