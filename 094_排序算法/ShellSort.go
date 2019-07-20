package main

//ShellSort 希尔排序(缩小增量排序)
func ShellSort(arr []int) {
	//获取增量 increment := len(arr)/2, len(arr)/2/2, len(arr)/2/2/2...
	for inc := len(arr) / 2; inc > 0; inc /= 2 {
		//获取比较的增量元素下标
		for i := inc; i < len(arr); i++ {
			temp := arr[i]
			for j := i - inc; j > 0; j -= inc {
				if temp < arr[j] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				} else {
					break
				}
			}
		}
	}
}
