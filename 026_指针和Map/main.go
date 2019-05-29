package main

import "fmt"

func main() {
	/*
		指针和Map
	*/

	a, b, c := 10, 20, 30

	//定义指针map
	m := make(map[int]*int)
	m[1] = &a
	m[2] = &b
	m[3] = &c

	//范围遍历打印变量的值
	for k, _ := range m {
		fmt.Println(*m[k])
	}

	//通过map修变量的值
	*m[1] = 100
	*m[2] = 200
	*m[3] = 300
	fmt.Println(a, b, c)

	//定义指针数组map
	m2 := map[int][3]*int{11: [3]*int{&a}, 12: [3]*int{&b, &c}}

	//范围遍历打印变量的值
	for _, v := range m2 {
		for i := 0; i < len(v); i++ {
			if v[i] != nil { //空指针不能操作
				fmt.Println(*v[i])
			}
		}
	}

	//通过map修改变量c的值
	*m2[12][1] = 999
	fmt.Println(c, "")
}
