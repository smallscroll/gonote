package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		定时器 Timer
		Timer 是一个结构体，包含了接收通道数据的管道成员C
	*/
	fmt.Println("now:", time.Now())

	//创建Timer结构体t1，系统会在2s后向通道发送系统当前时间，返回一个Timer指针
	//NewTimer创建一个Timer，它会在最少过去时间段d后到期，向其自身的C字段发送当时的时间。
	t1 := time.NewTimer(time.Second * 2)

	//定时器重设
	t1.Reset(time.Second * 4)

	//结构体t1.C接收Timer通道的数据（两步实现同步）
	retT := <-t1.C
	fmt.Println("t1:", retT)

	/*
		延迟定时!!!
	*/

	//time.After()：系统延迟发送系统当前时间到通道，并同时返回一个接收通道（一步实现同步）
	//After会在另一线程经过时间段d后向返回值发送当时的时间。等价于NewTimer(d).C
	fmt.Println("t2", <-time.After(time.Second*2))

	/*
		周期定时 Ticker
		周期定时不支持定时器重设
		Ticker保管一个通道，并每隔一段时间向其传递"tick"。
	*/

	//系统周期性(每2s)写入系统当前时间，返回一个Ticker指针
	//NewTicker返回一个新的Ticker，该Ticker包含一个通道字段，并会每隔时间段d就向该通道发送当时的时间。它会调整时间间隔或者丢弃tick信息以适应反应慢的接收者。如果d<=0会panic。关闭该Ticker可以释放相关资源。
	ti1 := time.NewTicker(time.Second * 2)

	i := 0
	for {
		fmt.Println("ti1", <-ti1.C) //Ticker指针接收通道数据
		i++
		if i == 3 {
			ti1.Stop() //停止定时器：3个周期后停止Ticker
			break
		}
	}
}
