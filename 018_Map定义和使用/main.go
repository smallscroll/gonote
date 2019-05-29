package main

import (
	"fmt"
)

func test(m map[int]string) {
	m[301] = "哈哈"
}

func main() {
	/*
		Map的定义和使用
	*/

	//定义map
	m := make(map[int]string)
	//添加数据
	m[101] = "张三"
	m[109] = "李四"
	//map的存储是无序的
	for k, v := range m {
		fmt.Println(k, v)
	}

	//定义并初始化map：键不能重复
	m2 := map[int]string{201: "王五", 202: "田七"}
	fmt.Println(m2)

	//通过键找到值
	fmt.Println(m2[202])
	//map名本身是一个地址
	fmt.Printf("%p\n", m2)
	//len()函数返回键值对个数
	fmt.Println(len(m2))

	//通过值找到键
	for k, _ := range m2 {
		if m2[k] == "田七" {
			fmt.Println(k)
		}
	}

	//判断键是否存在
	value, ok := m2[205]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("不存在")
	}

	//删除键值对(Map,键)
	delete(m2, 202)
	fmt.Println(m2)

	//Map作为函数传递：引用传递，形参可以改变实参的值
	test(m2)
	fmt.Println(m2)
}
