package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//创建随机数种子
	rand.Seed(time.Now().UnixNano())
	//生成1-100的随机数
	value := rand.Intn(100) + 1
	fmt.Println("猜一个1-100之间的数字：")
	num := 0
	//将输入的数字和随机数比较
	for {
		fmt.Scan(&num)
		if num > value {
			fmt.Println("太大了")
		} else if num < value {
			fmt.Println("太小了")
		} else {
			fmt.Println("猜对了(^.^)")
			break
		}
	}
}
