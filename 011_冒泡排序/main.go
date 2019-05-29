package main

import "fmt"

func main() {

	/*
		冒泡排序
	*/

	paopao := []int{6, 8, 3, 4, 2, 5, 9, 10, 1, 7}
	fmt.Println("原始泡泡：", paopao)
	//外层循环控制行，执行1次
	for i := 0; i < len(paopao)-1; i++ {
		//内层循环控制列，执行1周
		for j := 0; j < len(paopao)-1-i; j++ {
			//相邻两个元素比较，大于：降序，小于：升序
			if paopao[j] > paopao[j+1] {
				//交换相邻两个元素的数据
				paopao[j], paopao[j+1] = paopao[j+1], paopao[j]
			}
		}
	}
	fmt.Println("降序泡泡：", paopao)

}
