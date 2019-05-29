package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		select 实现超时
	*/

	//创建用于数据通信的channel
	ch := make(chan int)

	quit := make(chan bool)

	go func() { //子go程使用select监听channel数据流动
		for {
			select {
			case num := <-ch: //监听ch是否有数据流出
				fmt.Println("data:", num)
			case <-time.After(time.Second * 5): //阻塞状态，5秒后才会产生数据流动（向通道发送时间并接收），所有分支阻塞后，select监测到此分支数据流动后便执行该分支语句
				fmt.Println("超时退出")
				quit <- true //向退出通道发送数据
				return
			}
		}
	}()

	//主go程向channel发送数据
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(time.Second / 2)
	}

	<-quit //接收退出通道数据
}
