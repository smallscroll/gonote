package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 7, 1, 0, 3, 9, 8, 2, 5, 6}

	//冒泡排序
	//BubbleSort(arr)
	//BubbleSortBeta(arr)

	//选择排序
	//SelectSort(arr)

	//插入排序
	//InsertSort(arr)
	//InsertSortBeta(arr)

	//计数排序
	//CountingSort()

	//希尔排序
	//ShellSort(arr)

	//堆排序
	//HeapInit(arr)
	HeapSortLoop(arr)

	fmt.Println(arr)

	//二分查找
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	index := BinarySearch(slice, 9)
	fmt.Println(index)

}
