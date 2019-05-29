package main

import "fmt"

func main() {
	/*
		多级指针
	*/
	a := 10

	//一级指针：存储变量地址
	p := &a
	//*p 指向变量的值
	fmt.Println(*p)

	//二级指针：存储一级指针的地址
	var pp **int = &p

	//通过二级指针修改变量的值
	**pp = 100
	fmt.Println(a)

	//定义三级指针：存储二级指针的地址
	ppp := &pp
	//*ppp 指向二级指针的值
	//**ppp 指向一级指针的值
	//***ppp 指向变量的值
	***ppp = 1000
	fmt.Println(a)

}
