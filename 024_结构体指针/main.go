package main

import "fmt"

type Student struct {
	id   int
	name string
	addr string
}

func main() {
	/*
		结构体指针
	*/

	stu := Student{1001, "雷神", "阿斯加德"}

	var p *Student = &stu
	fmt.Printf("%p\n", p)
	//通过指针操作结构体的值
	(*p).name = "洛基"
	fmt.Println(stu)
	//可直接使用指针变量操作结构体的值
	p.name = "奥丁"
	fmt.Println(stu)

	//new()创建内存空间来存储结构体
	p2 := new(Student)
	p2.id = 1005
	p2.name = "蚁人"
	p2.addr = "纽约"
	fmt.Println(p2)
	fmt.Println(*p2)
	fmt.Printf("%T", p2)

}
