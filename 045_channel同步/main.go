package main

import (
	"fmt"
	"time"
)

/*
	channel 通道特性：

	• 通道中的数据只能单向流动
	• 通道中的数据只能接收一次，不能反复
	• 为保证数据在通道中正确流动，通道两端必须同时在线
	• 默认情况下，通道接收和发送数据都是阻塞的，除非另一端已经准备好
	  （一端发送，一端如果么接收的话该goroutine将处于阻塞状态）


	channel <- value      //发送value到channel
	<-channel             //接收并将其丢弃
    x := <-channel        //从channel中接收数据，并赋值给x
    x, ok := <-channel    //接收数据，同时检查通道是否已关闭或者是否为空

*/

//创建通道，数据类型设置为string
var ch = make(chan string)

//利用通道特性实现goroutine同步，使per1先打印，per2后打印

//封装打印机
func printer(str string) {
	for _, chr := range str {
		fmt.Printf("%c", chr)
		time.Sleep(time.Second / 10)
	}
}

//封装打印机使用者
func per1() {

	printer("hello")

	//给通道发送数据（如果另一端没有接收，该goroutine则会阻塞）
	ch <- "--over--"

}

func per2() {

	//接收通道数据（如果通道内没数据，该goroutine则会阻塞）
	s := <-ch

	fmt.Println(s)

	printer("world")

}

func main() {

	go per1()
	go per2()

	time.Sleep(time.Second * 3)
	// for { //防止主goroutine提前退出
	// 	runtime.GC()
	// }

}
