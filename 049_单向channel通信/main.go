package main

import (
	"fmt"
	"time"
)

/*
	双向channel:
	var ch chan int (由于没有make，该通道不能r/w)
	双向可以转换为单向channel

	单向channel:
	var chr <-chan int (单向r通道)
	var chw chan<- int (单向w通道)
	单向不可以转换为双向channel
*/

//接收函数 --读操作（单向读channel作为函数参数）
func recv(in <-chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("recv:", <-in) //in为单向读channel，只能接收数据，不能发送数据
		time.Sleep(time.Second)
	}

}

//发送函数 --写操作（单向写channel作为函数参数）
func send(out chan<- int) {
	for i := 0; i < 5; i++ {
		out <- i //out为单向写channel，只能发送数据，不能接收数据
	}

}

func main() {

	//定义双向channel
	ch := make(chan int)

	go send(ch) //ch引用传递给out

	recv(ch) //ch引用传递给in

}
