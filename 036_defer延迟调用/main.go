package main

import "fmt"

/*
	defer延迟调用
*/

func main() {

	/*
		defer延迟调用
	*/

	//defer函数入栈后并不马上调用，在出栈时按出栈顺序依次调用
	defer fmt.Println("hello world")
	defer fmt.Println("hahah！")
	fmt.Println("yayayaya~~")
	//
	a := 10
	b := 20
	defer func(a int, b int) {
		fmt.Println(a, b) //10 20
	}(a, b) //传入值为10 20
	a = 100
	b = 200
	fmt.Println(a, b) //100 200

}
