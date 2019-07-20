package main

import "fmt"

//BubbleSort 冒泡排序
func BubbleSort(arr []int) {
	count := 0

	//外层控制行
	for i := 0; i < len(arr)-1; i++ {
		//内层控制列
		for j := 0; j < len(arr)-1-i; j++ {
			count++
			//比较相邻数据大小
			if arr[j] > arr[j+1] {
				//满足条件则进行交换
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Printf("冒泡排序循环%d次\n", count)
}

//BubbleSortBeta 冒泡排序升级版
func BubbleSortBeta(arr []int) {
	count := 0

	flag := false
	//外层控制行
	for i := 0; i < len(arr)-1; i++ {
		//内层控制列
		for j := 0; j < len(arr)-1-i; j++ {
			count++
			//比较相邻数据大小
			if arr[j] > arr[j+1] {
				//满足条件则进行交换
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true //有数据交换时
			}
		}
		if !flag { //没有数据需要交换时
			fmt.Printf("冒泡排序升级版循环%d次\n", count)
			return
		}
		flag = false
	}
}
