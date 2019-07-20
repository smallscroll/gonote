package main

//SelectSort 选择排序
func SelectSort(arr []int) {
	//外层控制行
	for i := 0; i < len(arr)-1; i++ {
		//记录最大值下标的变量(初始下标为0)
		index := 0
		//内层控制列
		for j := 1; j < len(arr)-i; j++ {
			if arr[j] > arr[index] { //比较大小
				index = j //更新记录最大值下标的变量
			}
		}
		//将index对应的数据(本次循环比较的最大值)和最后一个数据进行交换
		arr[index], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[index]
	}
}
