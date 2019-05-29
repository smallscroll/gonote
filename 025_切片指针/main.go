package main

import "fmt"

//切片指针作为函数参数
func test(slice *[]int) {
	*slice = append(*slice, 4, 5, 6)
}

func main() {
	/*
		切片指针
	*/
	slice := []int{1, 2, 3}

	//定义指向切片的指针
	var p *[]int = &slice
	//切片指针变量需要加*才能进行切片取值
	fmt.Println((*p)[2])

	//切片本身存储3个值：数据地址（非切片地址）、长度、容量
	//虽然切片作为函数参数为引用传递（数据地址），但是：
	//切片传递到函数内追加数据时必须使用切片地址进行传递
	test(&slice)
	fmt.Println(slice)

}
