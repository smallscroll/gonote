package main

import "fmt"

func main() {

	/*
		select

		select的case语句中必须是一个I/O操作
		select同时监听所有case分支，监听不到数据流动（当前case不满足条件）的分支会阻塞
		多个case同时满足条件时，select随机选择一个执行
		一次selcet监听只能执行一个case分支，因此select通常在for循环中使用
		全部case阻塞之后会执行default分支，为防止忙轮询，for循环中的selcet通常省略default

		• 使用select的go程与其他go程之间采用的是【异步】通信方式

	*/

	ch := make(chan int)

	quit := make(chan bool)

	go func() { //子go程使用select监听channel数据流动

		for {
			select {
			case num := <-ch: //监听ch是否有数据流出，有则打印数据
				fmt.Println("data:", num)
			case <-quit:
				return // 等同于 runtime.Goexit() //退出当前goroutine
			}
		}
	}()

	//主go程向channel发送数据
	for i := 0; i < 10; i++ {
		ch <- i
	}
	quit <- true
	fmt.Println("finish")
}
