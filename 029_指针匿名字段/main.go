package main

import "fmt"

/*
	结构体不允许相互嵌套，但可以嵌套结构体指针
	指针在64位操作系统的内存中占6个字节大小
*/

type Person struct {
	name string
	sex  string
	addr string
}

type Student struct {
	*Person //指针类型匿名字段
	score   int
}

func main() {

	//对象初始化（取出地址给指针）
	stu := Student{&Person{"小明", "男", "北京"}, 98}
	//（*stu.Person).name = "小李"
	stu.name = "小李"
	fmt.Println(*stu.Person)

	//定义指针变量
	//stu2.Person的默认值为nil，属于空指针引用，无法操作
	var stu2 Student
	//为指针变量创建内存空间
	stu2.Person = new(Person)

	stu2.name = "Lucy"
	stu2.sex = "Lady"
	stu2.addr = "Earth"
	fmt.Println(*stu2.Person)
}
