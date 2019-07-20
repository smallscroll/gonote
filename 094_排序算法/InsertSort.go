package main

//InsertSort 插入排序
func InsertSort(arr []int) {
	//将数据分组为有序组和无序组
	//把arr的第一个元素作为有序组
	//第二个元素开始作为无序组
	for i := 1; i < len(arr); i++ {
		//将无序组数据与前一个有序组数据比较，小于有序组数据则插入到其前面
		if arr[i] < arr[i-1] {
			j := i - 1     //有序组的相邻数据的下标
			temp := arr[i] //无序组的第一个数据
			//在有序组内依次向前比较：
			for j >= 0 && arr[j] > temp {
				arr[j+1] = arr[j] //向后赋值，保留待插入的数据位置
				j--               //有序组数据下标递减
			}
			arr[j+1] = temp
		}
	}
}

//InsertSortBeta 插入排序升级版
func InsertSortBeta(arr []int) {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] { //如果无序组小于有序组数据
			//将该无序组数据插入到有序组中
			for j := i; j > 0; j-- { //有序组内进行冒泡
				if arr[j] < arr[j-1] { //比较相邻元素，小的前移
					arr[j], arr[j-1] = arr[j-1], arr[j]
				}
			}
		}
	}
}
