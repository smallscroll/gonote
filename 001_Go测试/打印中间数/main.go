package main

import (
	"fmt"
)

//寻找两个数组的中间数

func main() {
	a := []int{1, 2, 0, 3, 5}
	b := []int{8, 9, 3, 4, 6, 2, 7, 6}

	//拼接
	a = append(a, b[:]...)

	fmt.Println(a)

	//去重
	m := make(map[int]int)
	s := make([]int, 0)
	for i := 0; i < len(a); i++ {
		m[a[i]]++
	}
	for k := range m {
		s = append(s, k)
	}

	fmt.Println(s)

	//s = a //不去重

	//排序
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}

	fmt.Println(s)

	//打印中间数
	if len(s)%2 != 0 {
		fmt.Println(s[len(s)/2])
	} else {
		fmt.Println(s[len(s)/2-1], s[len(s)/2])
	}

}
