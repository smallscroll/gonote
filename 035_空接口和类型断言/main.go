package main

import "fmt"

func main() {
	/*
		空接口定义和使用
	*/

	//定义空接口（万能类型）
	var i interface{}

	//空接口可以储存任意类型的值
	i = 10
	fmt.Println(i)

	/*
		空接口类型断言
	*/

	//定义空接口类型切片
	var slice []interface{}
	slice = append(slice, 1, 3.14, "hello")
	fmt.Println(slice)

	//通过类型断言获取数据类型和值
	//格式：空接口数据.(数据类型)
	//value, ok = element.(T)，value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型
	//如果element里面确实存储了T类型的数值，那么ok返回true，否则返回false
	value, ok := slice[2].(int)
	if ok { //if ok == ture
		fmt.Println(value)
	} else {
		fmt.Println("不是int类型")
	}

	//通过循环和类型断言打印空接口数据
	for _, v := range slice { //分别返回下标和对应的值
		if value, ok := v.(int); ok == true { //第一个返回的是值（接口变量本身），第二个返回的是判断结果（bool）
			fmt.Println("整型数据:", value)
		} else if value, ok := v.(float64); ok { //ok == true
			fmt.Println("浮点型数据:", value)
		} else if value, ok := v.(string); ok {
			fmt.Println("字符串类型:", value)
		}
	}

	//通过switch进行类型断言
	for i := 0; i < len(slice); i++ {
		switch slice[i].(type) { //返回值对应的类型
		case int:
			value := slice[i].(int) //返回类型对应的值
			fmt.Println("整型：", value)
		case float64:
			value := slice[i].(float64)
			fmt.Println("浮点型：", value)
		case string:
			value := slice[i].(string)
			fmt.Println("字符串：", value)

		}
	}

}
