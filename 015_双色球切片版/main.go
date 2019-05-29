package main

import (
	"fmt"
	"math/rand"
	"time"
)

//冒泡排序函数
func BubbleSort(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func main() {
	/*
		打印双色球
	*/

	//创建随机数种子
	rand.Seed(time.Now().UnixNano())
	//定义一个红球的空切片
	red := []int{}
	//循环生成6个1-33的随机数
	for i := 0; i < 6; i++ {
		v := rand.Intn(33) + 1
		//循环判断生成的随机数是否与之前生成的随机数重复，如果重复则重新生成
		for j := 0; j < len(red); j++ {
			if v == red[j] {
				v = rand.Intn(33) + 1
			}
		}
		//将随机数追加到切片
		red = append(red, v)
	}
	//将切片传递给冒泡排序函数进行重新排序，打印排序后的红球切片和一个1-16的篮球随机数
	BubbleSort(red)
	fmt.Println(red, rand.Intn(16)+1)

}
