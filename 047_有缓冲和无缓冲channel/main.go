package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		无缓冲channel: 有r没有w,r端go程堵塞，有w,没有r,w端堵塞。因此要求两端同时在线。
			chr := make(chan int)
			chr := make(chan int, 0)

		有缓冲channel: 自带缓冲区——channel可以根据缓冲区大小，存储适量的数据。
			chr := make(chan int, 3)

		len()统计缓冲区剩余元素个数
		cap()统计缓冲区容量
	*/

	//创建缓冲区容量为3的channel
	ch := make(chan int, 3)

	//获取有缓冲channel的len和cap
	fmt.Println("ding...", len(ch), cap(ch))

	go func() { //子goroutine
		for i := 0; i < 7; i++ {
			ch <- i

			//主go程睡眠，由于子go程缓冲容量为3，能缓冲3条数据库，因此阻塞前依然能打印三次
			//缓冲区容量存满之后如果另一端没有接收数据则会堵塞
			fmt.Println("dong...", len(ch), cap(ch))
		}
	}()

	time.Sleep(time.Second * 5)

	//主goroutine
	for i := 0; i < 7; i++ {
		fmt.Println(<-ch) //接收通道数据并打印
		fmt.Println("main...", len(ch), cap(ch))
	}

	/*

		同步通信：无缓冲channel
		当一个调用发出，如果没有得到结果，那个这个调用就不返回	——阻塞
		类似打电话：要求通信双方必须同时在线

		异步通信：有缓冲channel
		当一个调用发出，不等待调用结果，直接返回	——非阻塞
		类似发短信：信息发送方，发送完直接结束
	*/

}
