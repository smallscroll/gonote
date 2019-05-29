package main

import (
	"fmt"
	"math/rand"
	"time"
)

//冒泡排序
func Paoao(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func main() {

	/*
		双色球Map版
	*/
	//创建随机数种子
	rand.Seed(time.Now().UnixNano())
	//打印多注
	var num int
	fmt.Println("输入数量：")
	fmt.Scan(&num)
	for i := 0; i < num; i++ {
		//定义空map和红球切片
		m := make(map[int]int)
		var red []int
		//根据map键不能重复的特性随机红球
		for len(m) < 6 {
			m[rand.Intn(33)+1] = 0
		}
		//遍历键添加到红球切片
		for k, _ := range m {
			red = append(red, k)
		}
		//红球排序并打印
		Paoao(red)
		fmt.Print(red)
		//打印随机篮球
		fmt.Print(rand.Intn(16)+1, "\n")
	}
}
