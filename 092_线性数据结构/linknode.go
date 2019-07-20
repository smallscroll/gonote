package main

import (
	"fmt"
	"reflect"
)

//LinkNode 单向链表节点
type LinkNode struct {
	Data interface{} //数据域
	Next *LinkNode   //指针域
}

//Student 学生类
type Student struct {
	name string
	addr string
}

//Create 创建链表
func (node *LinkNode) Create(data ...interface{}) {
	if node == nil || data == nil {
		return
	}
	if len(data) == 0 {
		return
	}

	//创建头节点
	head := node

	//遍历data，依次取出数据来创建单向列表
	for _, v := range data {
		//创建新节点并初始化
		newNode := new(LinkNode)
		newNode.Data = v
		newNode.Next = nil

		node.Next = newNode //将新节点赋值到当前节点的指针域
		node = node.Next    //更新当前节点为新节点
	}

	//将node更新为头节点
	node = head
}

//Print 递归打印链表
func (node *LinkNode) Print() {
	if node == nil { //容错的同时作为递归出口
		return
	}
	if node.Data != nil {
		fmt.Print(node.Data, " ")
	}
	//递归调用本函数打印
	node.Next.Print()
}

//PrintLoop 循环打印链表
func (node *LinkNode) PrintLoop() {
	if node == nil {
		return
	}
	for node.Next != nil {
		node = node.Next //跳过头节点
		if node.Data != nil {
			fmt.Print(node.Data, " ")
		}
	}
	fmt.Println("")
}

//Length 获取链表长度(节点数量)
func (node *LinkNode) Length() int {
	if node == nil {
		return -1
	}
	//定义计数器
	i := 0
	//循环统计数据节点个数
	for node.Next != nil {
		node = node.Next //后移节点
		i++
	}
	return i
}

//InsertByHead 头插法插入节点
func (node *LinkNode) InsertByHead(data interface{}) {
	if node == nil || data == nil {
		return
	}
	//创建并初始化新节点
	newNode := new(LinkNode)
	newNode.Data = data
	newNode.Next = nil

	//将原节点的下一个节点作为新节点的下一个节点
	newNode.Next = node.Next

	//将新节点作为头节点的下一个节点
	node.Next = newNode

}

//InsertByTail 尾插法插入节点
func (node *LinkNode) InsertByTail(data interface{}) {
	if node == nil || data == nil {
		return
	}
	//创建并初始化新节点
	newNode := new(LinkNode)
	newNode.Data = data
	newNode.Next = nil

	//找到原链表的最后一个节点
	for node.Next != nil {
		node = node.Next
	}
	//将新节点作为尾节点
	node.Next = newNode

}

//InsertByIndex 按位置插入节点
func (node *LinkNode) InsertByIndex(index int, data interface{}) {
	if node == nil || data == nil {
		return
	}
	if index < 0 || index > node.Length() {
		return
	}
	//如果插入位置在尾部
	if index == node.Length() {
		node.InsertByTail(data)
		return
	}
	//创建并初始化新节点
	newNode := new(LinkNode)
	newNode.Data = data
	newNode.Next = nil

	//记录插入节点的前一个节点
	preNode := node

	//找到待插入位置
	for i := 0; i < index; i++ {
		preNode = node
		node = node.Next
	}
	//将新节点的下一节点指向node
	newNode.Next = node

	//将preNode的下一节点指向新节点
	preNode.Next = newNode
}

//DeleteByIndex 按位置删除节点
func (node *LinkNode) DeleteByIndex(index int) {
	if node == nil {
		return
	}
	if index < 0 || index > node.Length() {
		return
	}
	//index位置的前一个节点
	preNode := node

	//循环获取index所对应节点
	for i := 0; i < index; i++ {
		preNode = node   //index的前一节点
		node = node.Next //index对应的节点
	}

	//将index前一节点指向index后一个节点
	preNode.Next = node.Next

	//将删除的节点置空，驱使GC工作
	node.Data = nil
	node.Next = nil
	node = nil
}

//DeleteByData 按数据删除节点
func (node *LinkNode) DeleteByData(data interface{}) {
	if node == nil || data == nil {
		return
	}

	//方法1:
	//保存节点初始值（头节点）
	head := node
	i := 0
	//遍历链表，依次比较节点的数据域（数值和类型）
	for node.Next != nil {
		if reflect.DeepEqual(node.Data, data) && reflect.TypeOf(node.Data) == reflect.TypeOf(data) {
			head.DeleteByIndex(i) //按下标删除节点
			return                //终止循环
		}
		node = node.Next
		i++
	}

	// //方法2:
	// preNode := node //保存待删除节点的前一节点
	// for node.Next != nil {
	// 	preNode = node
	// 	node = node.Next //后移
	// 	if reflect.DeepEqual(node.Data, data) && reflect.TypeOf(node.Data) == reflect.TypeOf(data) {
	// 		preNode.Next = node.Next //前一个节点指向后一个节点
	// 		//置空启动GC
	// 		nide.Data = nil
	// 		node.Next = nil
	// 		node = nil
	// 		return
	// 	}
	// }
}

//SearchByData 按数据查找下标
func (node *LinkNode) SearchByData(data interface{}) int {
	if node == nil || data == nil {
		return -1
	}
	//计数器
	i := 0
	for node.Next != nil {
		i++
		node = node.Next
		if reflect.DeepEqual(node.Data, data) && reflect.TypeOf(node.Data) == reflect.TypeOf(data) {
			return i
		}
	}
	return -1
}

//Destroy 销毁链表
func (node *LinkNode) Destroy() {
	if node == nil {
		return
	}
	//递归销毁链表节点
	node.Next.Destroy()
	//置空节点，驱使GC工作
	node.Data = nil
	node.Next = nil
	node = nil
}
