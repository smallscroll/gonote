package main

import (
	"fmt"
)

func main() {
	/*
		slice：线性表
	*/

	//创建切片
	slice := new(Slice)
	slice.Create(5, 20, 1, 2, 3, 4, 5)
	fmt.Println("切片结构：", slice)

	//追加元素
	slice.Append(6, 7, 8, 9, 10)

	//打印切片数据
	slice.Print()

	fmt.Println("长度：", slice.Len)
	fmt.Println("容量：", slice.Cap)

	//根据下标获取元素值
	ret := slice.Get(5)
	fmt.Println("根据下标获取数据：", ret)

	//根据元素值获取下标
	idx := slice.Search(10)
	fmt.Println("根据数据获取下标：", idx)

	//删除切片元素
	slice.Delete(3)
	slice.Print()

	//通过下标插入元素到切片
	slice.Insert(3, 9)
	slice.Print()

	//销毁切片
	slice.Destroy()
	fmt.Println("slice已销毁")

	/*
		list：单向链表
	*/

	//创建结构体对象
	stu1 := Student{"jack", "北京"}
	stu2 := Student{"yoyo", "深圳"}
	stu3 := Student{"lily", "广州"}

	//创建练表
	list := new(LinkNode)
	list.Create(stu1, stu2, stu3)
	fmt.Println(list)

	//打印链表
	// list.Print()
	list.PrintLoop()

	//获取链表长度
	fmt.Println("链表长度：", list.Length())

	stu4 := Student{"tom", "上海"}
	//头插法插入节点
	// list.InsertByHead(stu4)
	//尾插入插入节点
	// list.InsertByTail(stu4)
	//按位置插入节点
	list.InsertByIndex(2, stu4)

	//按位置删除节点
	list.DeleteByIndex(3)
	//按数据删除节点
	list.DeleteByData(stu4)

	//按数据查找下标
	idx = list.SearchByData(stu1)
	fmt.Println(idx)

	//打印链表
	list.PrintLoop()

}
