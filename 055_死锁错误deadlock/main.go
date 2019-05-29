package main

import "fmt"

func main() {
	/*
		死锁：运行时的一种错误，不是锁
	*/

	/*

		死锁1：同一个go程，使用同一个channel，自己发送数据自己接收，导致死锁

	*/

	ch11 := make(chan int)
	ch11 <- 100 //导致main一直阻塞
	num := <-ch11
	fmt.Println(num)

	/*

		死锁2: 发送接收位于两个go程，但是在通道发送/接收数据之后后才创建go程，导致前一个发送/接收一直阻塞，后面的go程无法开始

	*/

	ch21 := make(chan int)
	ch21 <- 200 //导致main一直阻塞
	go func() {
		num := <-ch21
		fmt.Println(num)
	}()

	/*

		死锁3（互相死锁）: 两个go程，两个channel，go程1接收channel1的数据写入channel2，go程2接收channel2的数据写入channel1，导致互相死锁

	*/

	ch31 := make(chan int)
	ch32 := make(chan int)

	go func() {
		for {
			select {
			case num := <-ch31: //...
				fmt.Println(num)
				ch32 <- 300 //相互阻塞
			}
		}
	}()

	for {
		select {
		case num := <-ch32: //相互阻塞
			fmt.Println(num)
			ch31 <- 400 //...
		}
	}

	/*

		死锁4: channel和互斥锁、读写锁混用可能会导致互锁

	*/

}
