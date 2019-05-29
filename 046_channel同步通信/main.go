package main

import "fmt"

func main() {

	//创建通信channel
	ch := make(chan int)

	//创建用于协调stdout的channel
	ch2 := make(chan bool)

	//子go程
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i             //将i传送到ch(一共发送了10次)
			fmt.Println("o", i) //使用stdout进行输出（主go程和子go程都使用stdout,因此需要进行同步协调）
			ch2 <- true         //再次发送数据到ch2进行协调通信
		}
	}()

	//主go程
	for i := 0; i < 10; i++ {

		num := <-ch           //接收ch数据并赋值给num(必须接收10次)
		<-ch2                 //ch2接收数据并丢弃（保证通道畅通，实现同步协调）
		fmt.Println("m", num) //使用stdout进行输出

	}

}
