package main

import "fmt"

/*
	回调函数（数据类型）：
		用户自定义一个函数，不直接在程序中显式调用，当特定条件满足时，由函数指针进行调用，或由系统自动调用。

*/

//定义函数指针(数据类型)：指向 func(x int, y bool) int 函数原型
type FUNCP func(x int, y bool) int

//主调函数（函数指针作为参数，其他两个参数为函数指针调用的函数所需参数）
func useCallback(x int, y bool, p FUNCP) int {
	return p(x, y)
}

//回调函数1(自定义)
func addOne(x int, y bool) int {
	if y == true {
		x += 1
	}
	return x
}

//回调函数2(自定义)
func subTen(x int, y bool) int {
	if y == true {
		x -= 10
	}
	return x
}

func main() {
	//定义一个函数指针类型的变量
	var p FUNCP
	p = addOne //将回调函数赋值给函数指针变量

	//通过函数指针使用回调函数
	result := useCallback(10, true, p)
	fmt.Println(result)

	// //直接使用回调函数
	// result = useCallback(15, true, subTen)
	// fmt.Println(result)

}
