package main

import "fmt"

/*
	方法定义和使用
*/

type INT int

//方法：为接收着（INT类型的对象）绑定的函数，为系统类型绑定函数时必须指定别名
//func (方法接收者)方法名(方法参数列表)返回值列表 {...}
//方法接收者即作为传递过来的第一个参数
func (a INT) add(b INT) {
	sum := a + b
	fmt.Println(sum)
}

/*
	实例1
*/

//父类
type Person struct {
	name string
	sex  string
	addr string
}

//子类
type Student struct {
	Person
	score int
}

//父类方法
func (per *Person) PrintInfo() {
	fmt.Printf("展示名片：姓名：%s，性别：%s，地址：%s。\n", per.name, per.sex, per.addr)
}

//子类方法
//方法接收者类型加*号为引用传递（指针变量），不加*号为值传递（普通变量）
func (stu *Student) EditInfo() {
	stu.name = "小李"
	fmt.Printf("大家好，我改了名字，现在叫%s。\n", stu.name)

}

//方法重写
func (stu *Student) PrintInfo() {
	fmt.Printf("资料保密\n")
}

func main() {

	//创建INT类型的对象
	var a INT = 10
	var b INT = 20

	//为对象绑定方法
	a.add(b)
	b.add(a)

	/*
		实例1
	*/

	//创建子类对象
	stu := Student{Person{"小明", "男", "北京"}, 98}

	fmt.Println(stu)

	//调用方法（普通变量和指针变量会自动转换，可省略*号）
	//(*stu).EditInfo()
	stu.EditInfo()

	//子类调用父类的方法
	stu.PrintInfo()

	//初始化父类并调用方法
	per := Person{"Lucy", "Lady", "Earth"}
	per.PrintInfo()
}
