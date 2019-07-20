package main

import (
	"fmt"
)

//StackNode 链式栈节点
type StackNode struct {
	Data interface{}
	Next *StackNode
}

//CreateStack 创建链式栈
func CreateStack(data ...interface{}) *StackNode {
	if data == nil {
		return nil
	}
	if len(data) == 0 {
		return nil
	}
	//定义记录下一节点的变量
	var nextNode *StackNode

	//创建链式栈对象
	stack := new(StackNode)

	//创建节点，按栈形结构组织数据
	for _, v := range data {
		newNode := new(StackNode)
		newNode.Data = v
		newNode.Next = nil
		//将新节点保存为stack
		stack = newNode
		//当前节点的下一个节点
		stack.Next = nextNode
		//新节点产生时,nextNode记录当前节点的下一节点
		nextNode = stack
	}

	//返回一个链式栈(头节点)
	return stack
}

//StackPrint 打印链式栈
func StackPrint(s *StackNode) {
	if s == nil {
		return
	}
	for s != nil {
		fmt.Print(s.Data, "")
		s = s.Next //后移s
	}
}

//StackLength 获取链式栈长度
func StackLength(s *StackNode) int {
	if s == nil {
		return -1
	}
	i := 0
	for s != nil {
		i++
		s = s.Next
	}
	return i
}

//StackPush 入栈(数据压栈) 传入旧节点和数据，返回一个新节点
func StackPush(s *StackNode, data interface{}) *StackNode {
	if s == nil {
		return nil
	}
	if data == nil {
		return s
	}
	//创建新节点
	newNode := new(StackNode)
	newNode.Data = data
	newNode.Next = s //将新节点放到最前面（新节点的next为原链表）

	return newNode //新节点作为链表的头
}

//StackPop 出栈(数据弹栈) 直接返回头节点的下一个节点，GC自动销毁
func StackPop(s *StackNode) *StackNode {
	if s == nil {
		return nil
	}
	//销毁当前节点
	//...
	return s.Next
}

//StackClear 链式栈清空
func StackClear(s *StackNode) {
	if s == nil { //递归出口
		return
	}
	StackClear(s.Next)
	s.Data = nil
	s.Next = nil
	s = nil
}
