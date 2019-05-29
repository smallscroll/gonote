package main

import "fmt"

/*
	递归函数
	在函数内部调用本函数称为递归
	在递归函数中如果没有出口会导致栈区溢出
*/

//计算n的阶乘
var sum int = 1

func test(n int) {
	if n == 1 {
		return
	}
	sum *= n //5*4*3*2
	test(n - 1)
}

func main() {
	test(5)
	fmt.Println(sum)
}
