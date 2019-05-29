package main

import "fmt"

/*
	recover错误拦截
*/

func test() {
	//在错误出现之前进行拦截
	defer func() {
		// //从panic状态中恢复并重新获得流程控制权
		//recover()

		//recover 返回值为interface类型，存储为panic调用时的错误信息
		//获取印错误信息
		result := recover()
		if result != nil {
			fmt.Println(result) //打印错误信息
		}
	}() // 匿名函数

	var p *int
	*p = 123 //panic异常
	fmt.Println("test")
}

func main() {

	test()
	fmt.Println("hello world")
}
