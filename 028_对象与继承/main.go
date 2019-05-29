package main

import "fmt"

//父类
type Person struct {
	name string
	sex  string
	addr string
}

//子类 将结构体理解为对象的属性（类）
type Student struct {
	Person     //匿名字段：实现继承关系，也可用实名字段：per Person，操作属性时用：stu.per.name
	score  int //如果子类和父类存在同名字段，操作对象属性时采用就近原则
}

//子类
type Teacher struct {
	Person
	subject string
}

//描述对象的行为（方法）
func (s Student) Sayhi() {
	fmt.Printf("大家好，我叫%s，我来自%s。\n", s.name, s.addr)
}

func main() {

	//创建并初始化对象：描述对象的属性
	per := Person{"Lucy", "lady", "Earth"}
	//包含父类属性的对象
	stu := Student{Person{"小明", "男", "北京"}, 98}
	tea := Teacher{Person{"法师", "男", "上海"}, "区块链"}
	fmt.Println(per)
	fmt.Println(stu)
	fmt.Println(tea)
	//操作对象属性，也可使用：stu.Person.name（也可操作父类中的同名字段属性）
	stu.name = "小张"
	//对象.方法
	stu.Sayhi()
}
