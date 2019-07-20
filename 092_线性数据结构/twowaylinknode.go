package main

import (
	"fmt"
)

//TowWayLinkNode 双向链表节点
type TowWayLinkNode struct {
	Data interface{} //数据域
	Prev *TowWayLinkNode
	Next *TowWayLinkNode //指针域
}

//Create 创建双向链表
func (node *TowWayLinkNode) Create(data ...interface{}) {
	if node == nil || data == nil {
		return
	}
	if len(data) == 0 {
		return
	}

	//创建头节点
	head := node

	//遍历data，依次取出数据来创建数据节点
	for _, v := range data {
		//创建新节点并初始化
		newNode := new(TowWayLinkNode)
		newNode.Data = v
		newNode.Prev = node
		newNode.Next = nil

		node.Next = newNode //将新节点赋值到当前节点的指针域
		node = node.Next    //更新当前节点为新节点
	}

	//将node更新为头节点
	node = head
}

//Print 递归正向打印双向链表
func (node *TowWayLinkNode) Print() {
	if node == nil { //容错的同时作为递归出口
		return
	}
	if node.Data != nil {
		fmt.Print(node.Data, " ")
	}
	//递归调用本函数打印
	node.Next.Print()
}

//PrintLoop 循环反向打印双向链表
func (node *TowWayLinkNode) PrintLoop() {
	if node == nil {
		return
	}
	//找到链表的尾节点并保存
	for node.Next != nil {
		node = node.Next
	}
	//倒序打印链表数据
	for node.Prev != nil {
		if node.Data != nil {
			fmt.Print(node.Data, " ")
		}
		node = node.Prev //依次前移节点
	}
}

//Length 获取双向链表长度(节点数量)
func (node *TowWayLinkNode) Length() int {
	if node == nil {
		return -1
	}
	//定义计数器
	i := 0
	//循环统计数据节点个数
	for node.Next != nil {
		i++
		node = node.Next //后移节点
	}
	return i
}

//InsertByIndex 按位置插入双向链表节点
func (node *TowWayLinkNode) InsertByIndex(index int, data interface{}) {
	if node == nil || data == nil {
		return
	}
	if index < 0 || index > node.Length() {
		return
	}
	//记录插入节点的前一个节点
	preNode := node

	//找到index对应的节点并保存
	for i := 0; i < index; i++ {
		preNode = node
		node = node.Next
		//循环结束我node指向index对应的节点，preNode指向前一个节点
	}

	//创建并初始化新节点
	newNode := new(TowWayLinkNode)
	newNode.Data = data
	newNode.Prev = preNode
	newNode.Next = node

	//node的prev指向新节点
	node.Prev = newNode

	//preNode的next为新节点
	preNode.Next = newNode

}

//DeleteByIndex 按位置删除双向列表节点
func (node *TowWayLinkNode) DeleteByIndex(index int) {
	if node == nil {
		return
	}
	if index < 0 || index > node.Length() {
		return
	}

	l := node.Length()

	//记录index位置对应的前一个节点
	preNode := node

	//循环移动获取index所对应节点
	for i := 0; i < index; i++ {
		preNode = node   //index的前一节点
		node = node.Next //index对应的节点
	}

	//如果index为最后一个节点
	if index == l {
		preNode.Next = nil
		node.Data = nil
		node.Prev = nil
		node.Next = nil
		node = nil
		return
	}

	//将index前一节点指向index后一个节点
	preNode.Next = node.Next

	//将node下一节点的prev指向前一个节点
	node.Next.Prev = preNode

	//将删除的节点置空，驱使GC工作
	node.Data = nil
	node.Prev = nil
	node.Next = nil
	node = nil
}

//Destroy 销毁双向链表
func (node *TowWayLinkNode) Destroy() {
	if node == nil {
		return
	}
	//递归销毁链表节点
	node.Next.Destroy()
	//置空节点，驱使GC工作
	node.Data = nil
	node.Prev = nil
	node.Next = nil
	node = nil
}
