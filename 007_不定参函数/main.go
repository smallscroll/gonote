package main

import "fmt"

//函数接收不定参数
func add(arr ...int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func main() {
	//向函数传递多个参数
	fmt.Println(add(1, 2, 3, 4, 5))

	//向函数传递数组
	//数组作为函数参数传递为值传递
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(add(arr[:]...))
	fmt.Println(add(arr[0:2]...))

}
