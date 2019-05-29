package main

import (
	"fmt"
	"time"
)

/*
	生产者、消费者模型

	生产者模块：向市场（公共缓冲区）产生数据
	消费者模块：从市场（公共缓冲区）消费数据
	市场（公共缓冲区）：缓存生产者产生的数据

	缓冲区存在的好处：
	• 解耦：降低生产者和消费者之间的耦合度
	• 并发：生产者、消费者采用异步通信，提高并发性
	• 缓存：当生产和消费效率不对等时，双方可以利用市场提高通信效率
*/

//生产者模块 --产生数据
func producer(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i
		time.Sleep(time.Second / 5)
	}
	close(p) //关闭通道

}

//消费者模块 --消费数据
func consumer(c <-chan int) {
	for num := range c {
		fmt.Println("consumer:", num)
	}

}

func main() {

	//定义双向有缓冲channel（模拟市场）
	ch := make(chan int, 5)

	go producer(ch) //ch引用传递给p

	consumer(ch) //ch引用传递给c

}
