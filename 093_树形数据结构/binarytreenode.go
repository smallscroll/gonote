package main

import (
	"fmt"
	"reflect"
)

//BinaryTreeNode 二叉树节点
type BinaryTreeNode struct {
	Data   interface{}
	LChild *BinaryTreeNode
	RChild *BinaryTreeNode
}

//Create 按图例创建二叉树
func (node /*根节点*/ *BinaryTreeNode) Create() {

	//创建二叉树子节点
	node1 := BinaryTreeNode{1, nil, nil}
	node2 := BinaryTreeNode{2, nil, nil}
	node3 := BinaryTreeNode{3, nil, nil}
	node4 := BinaryTreeNode{4, nil, nil}
	node5 := BinaryTreeNode{5, nil, nil}
	node6 := BinaryTreeNode{6, nil, nil}
	// node7 := BinaryTreeNode{7, nil, nil}

	node.Data = 0
	node.LChild = &node1
	node.RChild = &node2

	node1.LChild = &node3
	node1.RChild = &node4

	node2.LChild = &node5
	node2.RChild = &node6

	// node3.LChild = &node7
}

//PreOrder 打印二叉树：先(根)序遍历 - DLR(中-左-右)
func (node *BinaryTreeNode) PreOrder() {
	if node == nil { //递归出口
		return
	}
	//D 先打印data数据
	fmt.Print(node.Data, " ")
	//L 左子树递归调用本函数
	node.LChild.PreOrder()
	//R 右子树递归调用本函数
	node.RChild.PreOrder()
}

//MidOrder 打印二叉树：中(根)序遍历 - LDR(左-中-右)
func (node *BinaryTreeNode) MidOrder() {
	if node == nil { //递归出口
		return
	}
	//L 左子树递归调用本函数
	node.LChild.MidOrder()
	//D 打印data数据
	fmt.Print(node.Data, " ")
	//R 右子树递归调用本函数
	node.RChild.MidOrder()
}

//PostOrder 打印二叉树：后(根)序遍历 - LDR(左-右-中)
func (node *BinaryTreeNode) PostOrder() {
	if node == nil { //递归出口
		return
	}
	//L 左子树递归调用本函数
	node.LChild.PostOrder()
	//R 右子树递归调用本函数
	node.RChild.PostOrder()
	//D 打印data数据
	fmt.Print(node.Data, " ")
}

//TreeHeight 获取二叉树深度（高度）
func (node *BinaryTreeNode) TreeHeight() int {
	if node == nil { //容错同时作为递归出口
		return 0
	}
	//左子树递归进入
	lh := node.LChild.TreeHeight()
	//右子树递归进入
	rh := node.RChild.TreeHeight()
	//累加并比较左右子树的高度
	if lh > rh {
		lh++
		return lh
	}

	rh++
	return rh
}

//LeafNum 获取叶子节点的个数
func (node *BinaryTreeNode) LeafNum(num *int) {
	if node == nil {
		return
	}
	//判断是否为叶子节点
	if node.LChild == nil && node.RChild == nil {
		*num++
	}
	//左/右子树各自递归调用本函数
	node.LChild.LeafNum(num)
	node.RChild.LeafNum(num)
}

//Search 二叉树数据查找(判断数据是否存在)
func (node *BinaryTreeNode) Search(data interface{}) {
	if node == nil {
		return
	}
	//比较数据值和类型
	if reflect.DeepEqual(node.Data, data) && reflect.TypeOf(node.Data) == reflect.TypeOf(data) {
		fmt.Println("找到数据：", node.Data)
		return
	}
	//左/右子树各自递归调用本函数
	node.LChild.Search(data)
	node.RChild.Search(data)
}

//Destroy 二叉树销毁
func (node *BinaryTreeNode) Destroy() {
	if node == nil {
		return
	}

	node.LChild.Destroy()
	node.LChild = nil

	node.RChild.Destroy()
	node.RChild = nil

	node.Data = nil
	node = nil
}

//Reverse 二叉树反转(满二叉树)
func (node *BinaryTreeNode) Reverse() {
	if node == nil {
		return
	}
	//左子树和右子树交换
	node.LChild, node.RChild = node.RChild, node.LChild

	//左/右子树各自递归调用本函数
	node.LChild.Reverse()
	node.RChild.Reverse()
}

//Copy 二叉树拷贝
func (node *BinaryTreeNode) Copy() *BinaryTreeNode {
	if node == nil {
		return nil
	}
	//左子树自递归调用本函数
	lChild := node.LChild.Copy()
	//右子树自递归调用本函数
	rChild := node.RChild.Copy()

	//创建新节点并赋值
	newNode := new(BinaryTreeNode)
	newNode.Data = node.Data
	newNode.LChild = lChild
	newNode.RChild = rChild

	//返回新节点
	return newNode
}
