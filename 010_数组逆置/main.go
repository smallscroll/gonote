package main

import "fmt"

func main() {
	/*
		数组逆置

		//数组的元素个数为常量或常量表达式，定义后不能修改
	*/

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//起始元素的索引值
	start := 0
	//终止元素的索引值
	end := len(arr) - 1
	//开始循环，起始索引值必须小于终止索引值
	for start < end {
		//起始与终止的数据依次进行交换
		arr[start], arr[end] = arr[end], arr[start]
		//起始索引值递增
		start++
		//终止索引值递减
		end--
	}
	//打印结果
	fmt.Println(arr)

}
