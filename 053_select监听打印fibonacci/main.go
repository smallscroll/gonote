package main

import "fmt"

func fibonacci(chr <-chan int, quit chan<- string) {
	//循环监听
	for {
		select {
		case num := <-chr: //监听数据流出（读数据）
			fmt.Println(num)
		case quit <- "退出": //发送数据到退出通道
			return
		}
	}

}

func main() {

	/*	fibonacci
		斐波那契数列：
		1, 1, 2, 3, 5, 8, 13, 21, 34 ...
		x, y
		   x, y
			  x, y
			     x, y

	*/

	//创建用于数据传送的channel
	ch := make(chan int)
	//创建用于退出的channel
	quit := make(chan string)

	//创建子go程
	go fibonacci(ch, quit)

	//数据生成
	x, y := 1, 1
	for i := 0; i < 15; i++ {
		ch <- x //将每次循环新生成的x发送到channel
		x, y = y, x+y
	}

	//接收退出通道数据
	<-quit
}
