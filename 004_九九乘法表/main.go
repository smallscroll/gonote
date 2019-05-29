package main

import "fmt"

func main() {
	//循环嵌套规则：
	//外层循环执行1次，内层循环执行1周
	//外层循环控制行，内层循环控制列

	//九九乘法表
	/*
		1*1=1
		1*2=2 2*2=4
		1*3=4 2*3=6 3*3=9
		...
	*/
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, j*i)
		}
		fmt.Println()
	}

}
