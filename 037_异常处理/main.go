package main

import (
	"errors"
	"fmt"
)

/*
	error异常
*/
//主动捕获错误信息，返回指针类型默认值为nil
func test(a int, b int) (v int, err error) {
	if b == 0 {
		err = errors.New("除数不能为零")
		return
	}
	v = a / b
	return
}

/*
	panic异常
*/

func test1(i int) {
	var arr [3]int = [3]int{1, 2, 3}
	//当程序出现不可恢复的错误时，系统会调用panic来终止程序并打印错误信息
	fmt.Println(arr[i])
}

func main() {

	/*
		error异常
	*/

	v, err := test(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}

	/*
		panic异常
	*/
	//test1(10)

	fmt.Println("hello~1")

	//调用panic，终止程序执行
	panic("hello world")

	fmt.Println("hello~2")
	fmt.Println("hello~3")

}
