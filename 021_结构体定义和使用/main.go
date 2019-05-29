package main

import "fmt"

//定义结构体(结构体是一种数据类型，可以用来定义数据格式)
type Student struct {
	//结构体成员
	id   int
	name string
}

//定义结构体
type Students struct {
	id   int
	name string
	age  int
}

//将结构体切片元素按年龄排序
func Paopao(slice []Students) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j].age < slice[j+1].age {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

//定义结构体
type Hero struct {
	name string
	age  int
	addr string
}

func main() {
	//定义结构体变量
	var stu Student
	// //定义并初始化结构体变量
	// stu := Student{530324, "张三"}

	//为结构体成员赋值
	stu.id = 530324
	stu.name = "张三"
	fmt.Println(stu)
	fmt.Printf("%p\n", &stu)
	fmt.Println(&stu.id)
	fmt.Println(&stu.name)

	//结构体赋值(同类型的两个结构体变量可以相互赋值)
	// stu1 := stu
	var stu1 Student
	stu1 = stu
	fmt.Println(stu1)
	stu1.name = "李四"
	fmt.Println(stu1)
	fmt.Println(stu)

	/*
		结构体数组
	*/
	fmt.Println("--------------------")

	//定义一个包含多的元素的结构体数组
	var arr [3]Students

	//为结构体数组元素赋值
	arr[0].id = 101
	arr[0].name = "钢铁侠"
	arr[0].age = 50
	fmt.Println(arr)

	/*
		结构体切片
	*/

	//定义并初始化结构体切片
	var slice []Students = []Students{
		{1001, "钢铁侠", 50},
		{1002, "绿巨人", 40},
		{1003, "美国队长", 70},
	}
	//打印所有切片元素
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	fmt.Println("----")

	//添加结构体切片元素
	slice = append(slice, Students{1004, "雷神", 200})
	//将结构体切片按年龄排序
	Paopao(slice)
	//打印所有切片元素
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	/*
		结构体变量作为函数参数：值传递
		结构体数组作为函数参数：值传递
		结构体切片作为函数参数：引用传递
	*/

	/*
		结构体Map
	*/

	//定义并初始化结结构体Map
	hero := map[int]Hero{
		1001: Hero{"雷神", 200, "阿斯加德"},
		1002: Hero{"绿巨人", 40, "巴西"},
		1003: Hero{"蝙蝠侠", 50, "哥谭市"},
	}
	//数据添加
	hero[1004] = Hero{"惊奇队长", 30, "银河"}
	//数据删除
	delete(hero, 1002)
	//打印Map
	fmt.Println(hero)
}
