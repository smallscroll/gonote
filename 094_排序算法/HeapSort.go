package main

/*
	堆排序
*/

//HeapInit 初始化堆
func HeapInit(arr []int) {

	//将切片转成二叉树模型  实现大根堆
	length := len(arr)
	for i := length/2 - 1; i >= 0; i-- { // 4,3,2,1
		HeapSort(arr, i, length-1)
	}

	//根节点存储最大值
	for i := length - 1; i > 0; i-- {
		//如果只剩下根节点和跟节点下的左子节点
		if i == 1 && arr[0] <= arr[i] {
			break
		}
		//将根节点和叶子节点数据交换
		arr[0], arr[i] = arr[i], arr[0]
		HeapSort(arr, 0, i-1)
	}
}

//HeapSort 堆排序：递归法（获取堆中最大值放在根节点）
func HeapSort(arr []int, startNode int, maxNode int) {

	// 最大值放在根节点
	var max int

	// 定义左子节点和右子节点
	lChild := startNode*2 + 1
	rChild := lChild + 1
	// 子节点超过比较范围 跳出递归
	if lChild >= maxNode {
		return
	}
	// 左右比较  找到最大值
	if rChild <= maxNode && arr[rChild] > arr[lChild] {
		max = rChild
	} else {
		max = lChild
	}

	// 和跟节点比较
	if arr[max] <= arr[startNode] {
		return
	}

	// 交换数据
	arr[startNode], arr[max] = arr[max], arr[startNode]

	// 递归进行下次比较
	HeapSort(arr, max, maxNode)
}

/*
	循环法：
*/

//HeapSortLoop 大顶堆排序
func HeapSortLoop(arr []int) {
	length := len(arr) // 获取数组长度

	// 将切片，转成二叉树模型。 实现大根堆（最大堆、大顶堆）
	for i := length/2 - 1; i >= 0; i-- { // 第一个非叶子节点的序号为length/2-1 —— 3、2、1、0
		HeapAdjustDown(arr, i, length-1)
	}

	// 进行堆排序
	for i := length - 1; i > 0; i-- {
		// 堆顶元素和最后一个元素交换位置, 最后一个位置保存最大数
		arr[0], arr[i] = arr[i], arr[0]

		// 将 arr[0...i-1] 重新调整为最大堆
		HeapAdjustDown(arr, 0, i-1)
	}
}

//HeapAdjustDown ...
func HeapAdjustDown(arr []int, start int, end int) {

	temp := arr[start] // 保存当前结点
	i := 2*start + 1   // 该结点的“左孩子”在数组中的位置序号

	for i <= end {
		// 找左、右孩子中的最大，用 i 记录位置序号
		if i+1 <= end && arr[i+1] > arr[i] { // i+1 " 右孩子" 的位置序号
			i++ // 右孩子大, i 记录右孩子位置序号
		}
		// 如果符合大顶堆定义（左、右均 <= 父结点），则不用调整位置
		if arr[i] <= temp {
			break
		}
		// 最大子结点，替换掉其父结点
		arr[start] = arr[i]
		start = i
		// 从 上->下 找
		i = 2*start + 1
	}
	arr[start] = temp
}
