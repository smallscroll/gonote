package main

import (
	"fmt"
)

func main() {
	//创建二叉树
	tree := new(BinaryTreeNode)
	tree.Create()
	fmt.Println(tree)

	//先序遍历
	fmt.Print("先序遍历：")
	tree.PreOrder()
	//中序遍历
	fmt.Print("中序遍历：")
	tree.MidOrder()
	//后序遍历
	fmt.Print("后序遍历：")
	tree.PostOrder()

	fmt.Println("")

	//获取二叉树深度
	fmt.Println("树高：", tree.TreeHeight())

	//获取叶子节点个数
	num := 0
	tree.LeafNum(&num)
	fmt.Println("叶子节点个数：", num)

	//二叉树数据查找
	tree.Search(1)

	//二叉树反转
	tree.Reverse()
	tree.MidOrder()

	fmt.Println("")

	//二叉树拷贝
	newTree := tree.Copy()
	fmt.Println(newTree)
	newTree.MidOrder()

}
