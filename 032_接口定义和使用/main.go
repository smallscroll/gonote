package main

import "fmt"

//接口类型：描述一系列方法的集合，是一种规范和标准
type Personer interface {
	//在接口中定义函数的格式
	SayHi()
}

//person对象类
type Person struct {
	name string
	age  int
	addr string
}

//person方法实现
func (per *Person) SayHi() {
	fmt.Printf("大家好，我叫%s，我今年%d岁，我来自%s。\n", per.name, per.age, per.addr)
}

/*
	多态实现
*/

//多态：将接口类型作为函数参数进行封装（调用同一函数实现不同表现）
func Sayhi(p Personer) {
	//通过接口调用对象方法
	p.SayHi()
}

func main() {
	//实例化对象
	per := Person{"小明", 18, "北京"}

	//创建接口类型变量，接口可以接收任意类型数据
	var p Personer
	//将对象地址赋值给接口变量
	p = &per
	//通过接口调用对象方法（自动匹配对象类型）
	p.SayHi()

	/*
		多态实现
	*/

	//通过多态调用对象的方法
	Sayhi(&per)
}
