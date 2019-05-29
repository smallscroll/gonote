package main

import "fmt"

func main() {

	/*
		切片的定义
	*/

	//定义一个空切片
	var slice []int
	//空切片的长度为0，不能通过索引添加数据，须要为空切片追加数据
	slice = append(slice, 1, 2, 3)
	fmt.Println(slice)

	//定义指定长度为10的切片，容量默认等于长度10
	//slice2 := make([]int, 10)
	//定义指定长度为10，容量为20的切片，可根据长度索引添加数据
	slice2 := make([]int, 10, 20)
	slice2[1] = 11
	slice2[3] = 13
	fmt.Println(slice2)
	//计算切片的长度
	fmt.Println(len(slice2))
	//计算切片的容量
	fmt.Println(cap(slice2))
	//定义长度/容量的切片追加数据超过该容量后，切片容量将自动按照定义值的2倍扩充
	//未定义长度/容量的切片追加数据后按偶数倍递增自动扩容
	//切片长度超过1024后，扩容按照上一次的1/4容量扩容
	slice2 = append(slice2, 999, 888, 777, 666)
	fmt.Println(slice2)
	fmt.Println(cap(slice2))

	/*
		打印切片数据
	*/

	fmt.Println("打印切片数据：")

	//通过范围遍历打印切片数据，i和v分别接收索引和数据
	for i, v := range slice2 {
		fmt.Println(i, v)
	}
	//当范围遍历只有一个接收值时表示接收切片的索引（下标）
	for i := range slice2 {
		fmt.Println(i)
	}

	/*
		切片的截取
	*/

	fmt.Println("切片的截取：")

	//切片名即为地址，为另一个切片赋值后再修改的为同一个切片
	s2 := slice2
	slice2[0] = 111
	fmt.Println(s2)
	fmt.Println(slice2)

	//切片拷贝为两个不同的地址，重新赋值后不影响另一个切片
	s3 := make([]int, 20)
	copy(s3, slice2)
	s3[4] = 555
	fmt.Println(slice2)
	fmt.Println(s3)

	//切片截取是在原切片上操作，属于同一地址，修改后会影响原切片
	s4 := s3[0:3]
	fmt.Println(s4)
	s4[0] = 222
	fmt.Println(s3)
	//切片[起始位置：结束位置：最大容量] cap = 最大容量 - 起始位置
	s5 := s3[2:5:20]
	fmt.Println(cap(s5))

	/*
		切片数据间接删除（截取-追加）
	*/
	s6 := []int{10, 20, 30, 40, 50}
	s6 = append(s6[2:3], s6[len(s6)-1:]...)
	fmt.Println(s6)

	/*
		//切片名包含地址、长度、容量3个值
		//函数中切片作为地址（引用）传递，形参和实参指向相同地址
	*/
}
