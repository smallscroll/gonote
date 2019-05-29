package main

import "fmt"

//指针作为函数参数
func test(a *int, b *int) {
	*a, *b = *b, *a //去值并进行交换
}

//数组指针作为函数参数
func paopao(arr *[5]int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	/*
		指针的定义和使用
	*/
	var i int = 10
	fmt.Println(&i)

	//定义指针变量
	var p *int = &i
	//通过指针间接修改变量值
	*p = 11
	fmt.Println(i) //11

	// //空指针：指针变量指向内存地址空间为0（系统占用、禁止用户读写）
	// var p *int
	// fmt.Println(p) //默认值为nil
	// //野指针：指针变量指向内存的未知空间（禁止操作）

	/*
		开辟内存空间
	*/

	var p2 *int
	//开辟数据类型对应的内存空间并初始化
	p2 = new(int)
	*p2 = 100
	fmt.Println(p2)

	//将指针置空，回收内存空间（垃圾回收机制）
	p2 = nil
	//垃圾回收机制（gc）的两种方式:
	//1.标记清除、三色标记，使用在超过32KB大小的内存
	//2.引用计数，如果使用计数+1，如果没有使用计数-1，引用计数为0时被系统回收
	fmt.Println(p2)

	/*
		指针作为函数参数
	*/

	a := 10
	b := 20
	fmt.Println(a, b)
	//传递地址（引用传递）
	test(&a, &b)
	fmt.Println(a, b)

	/*
		数组指针
	*/

	arr := [3]int{1, 2, 3}
	p3 := &arr //var p3 *[]int = &arr
	fmt.Println(p3)
	//数组指针可直接使用指针变量取值，等同于(*p3)[2]
	fmt.Println(p3[2])

	//开辟堆空间时，可直接将指针当作数组名一样使用
	var p4 *[3]int = new([3]int)
	fmt.Println(p4)

	//数组指针作为函数参数为引用传递（地址）
	arr2 := [5]int{4, 2, 1, 5, 3}
	paopao(&arr2)
	fmt.Println(arr2)

	/*
		指针数组
	*/

	a2, b2, c2 := 21, 22, 23
	var arr3 [3]*int = [3]*int{&a2, &b2, &c2}
	//指针数组存储的为变量地址
	fmt.Println(arr3)
	//修改元素数据
	*arr3[0] = 301
	//打印指针数组元素的值
	for i := 0; i < len(arr3); i++ {
		fmt.Println(*arr3[i])
	}

	/*
		数组指针数组
	*/
	a3 := [3]int{1, 2, 3}
	b3 := [3]int{4, 5, 6}
	c3 := [3]int{7, 8, 9}
	//定义并初始化数组指针数组
	var arr4 [3]*[3]int = [3]*[3]int{&a3, &b3, &c3}
	//打印数组的值
	for i := 0; i < len(arr4); i++ {
		fmt.Println(arr4[i])
	}
	//打印数组元素的值
	for i := 0; i < len(arr4); i++ {
		for j := 0; j < len(arr4[i]); j++ {
			fmt.Print(arr4[i][j])
		}
		fmt.Println("")
	}

}
