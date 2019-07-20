package main

import (
	"fmt"
)

//LoopLinkNode 环形链表节点
type LoopLinkNode struct {
	Data interface{}   //数据域
	Next *LoopLinkNode //指针域
}

//Create 创建环形链表
func (node *LoopLinkNode) Create(data ...interface{}) {
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
		newNode := new(LoopLinkNode)
		newNode.Data = v
		newNode.Next = nil

		node.Next = newNode //将新节点赋值到当前节点的指针域
		node = node.Next    //更新当前节点为新节点
	}

	//将尾节点的next指向第一个数据节点，形成逻辑环形
	node.Next = head.Next

	//还原头节点
	node = head
}

//PrintLoop 循环打印环形链表
func (node *LoopLinkNode) PrintLoop() {
	if node == nil {
		return
	}

	//创建标记节点:第一个数据节点
	start := node.Next

	//打印链表数据
	for {
		node = node.Next
		if node.Data != nil {
			fmt.Print(node.Data, " ")
		}
		if start == node.Next {
			break
		}
	}
}

//Length 获取环形链表长度(节点数量)
func (node *LoopLinkNode) Length() int {
	if node == nil {
		return -1
	}
	//定义标记位记录第一个数据节点
	start := node.Next

	//定义计数器
	i := 0
	//循环统计数据节点个数
	for {
		node = node.Next //后移节点
		i++
		if start == node.Next {
			break
		}
	}
	return i
}

//InsertByIndex 按位置插入环形链表节点
func (node *LoopLinkNode) InsertByIndex(index int, data interface{}) {
	if node == nil || data == nil {
		return
	}
	if index < 0 || index > node.Length() {
		return
	}
	//定义标记位记录第一个数据节点
	start := node.Next

	//记录index的前一个节点
	preNode := node

	//找到index对应的节点并保存
	for i := 0; i < index; i++ {
		preNode = node
		node = node.Next
		//循环结束时node指向index对应的节点，preNode指向前一个节点
	}

	//创建并初始化新节点
	newNode := new(LoopLinkNode)
	newNode.Data = data
	newNode.Next = node

	//preNode的next为新节点
	preNode.Next = newNode

	//如果插入位置是1
	if index == 1 {
		//找到尾节点，保存到node
		for {
			if start == node.Next {
				break //找到尾节点并已保存在node中
			}
			node = node.Next
		}
		//将尾节点的Next指向新数据节点
		node.Next = newNode
	}

}

//DeleteByIndex 按位置删除环形列表节点
func (node *LoopLinkNode) DeleteByIndex(index int) {
	if node == nil {
		return
	}
	if index < 0 || index > node.Length() {
		return
	}

	//定义标记位记录第一个数据节点
	start := node.Next

	//记录index位置对应的前一个节点
	preNode := node

	//循环移动获取index所对应节点
	for i := 0; i < index; i++ {
		preNode = node   //index的前一节点
		node = node.Next //index对应的节点
	}

	//当index为1时
	if index == 1 {
		temp := node //尾节点临时变量
		for {
			if start == temp.Next {
				break //节点保存在temp中
			}
			temp = temp.Next
		}
		//将尾节点的next指向删除节点的下一节点
		temp.Next = node.Next
	}

	//将index前一节点指向index后一个节点
	preNode.Next = node.Next

	//将删除的节点置空，驱使GC工作
	node.Data = nil
	node.Next = nil
	node = nil
}

//Destroy 销毁环形链表
func (node *LoopLinkNode) Destroy() {
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
