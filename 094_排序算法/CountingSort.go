package main

import (
	"fmt"
	"math/rand"
	"time"
)

//CountingSort 计数排序（计数统计排序）
func CountingSort() {

	//随机生成10万个1000以内的随机数
	s := make([]int, 0)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100000; i++ {
		s = append(s, rand.Intn(1000))
	}

	//使用map统计1000以内的每个数出现的次数
	m := make(map[int]int)
	//遍历随机数切片，统计不同随机数的出现次数
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	//通过map输出所有数据进行排序
	for i := 0; i < 1000; i++ {
		//内层循环：将每个数出现多少次就打印多少次
		for j := 0; j < m[i]; j++ {
			fmt.Println(i, " ") //实现排序
		}
	}
}
