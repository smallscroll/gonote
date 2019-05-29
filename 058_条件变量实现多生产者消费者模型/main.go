package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	条件变量：sync.Cond类型代表了条件变量，条件变量要与锁（互斥锁，或者读写锁）一起使用，成员变量L代表与条件变量搭配使用的锁。
		type Cond struct{
			...
			L（锁：互斥/读写）
			...
		}

	Cond.Wait():
		释放已掌握的互斥锁相当于cond.L.Unlock()，阻塞等待条件变量满足
		当被唤醒，Wait()函数返回时，解除阻塞并重新获取互斥锁。相当于cond.L.Lock()

	Cond.Signal():
		唤醒另一端阻塞在条件变量上的一个对象

	Cond.Broadcast():
		唤醒另一端阻塞在条件变量上的所有对象


	条件变量使用流程：
		1.创建条件变量
		2.初始化条件变量使用的锁
		3.给条件变量的互斥锁加锁
		4.判断条件变量是否满足（使用for循环判断）
			调用wait()
		5.生产数据
		6.写入公共缓冲区
		7.唤醒对端
		8.给条件变量的互斥锁解锁


*/

var cond sync.Cond //创建全局条件变量

//生产者
func producer(p chan<- int, idx int) {
	for {

		cond.L.Lock()          //给条件变量的成员L进行加锁（和消费者是同一个锁）
		for len(p) == cap(p) { //判断市场满，等待消费者消费
			cond.Wait() //挂起当前goroutine,等待条件满足,被消费者唤醒（解锁并阻塞当前线程,在唤醒返回前加锁）
		}
		num := rand.Intn(1000) //生产数据
		p <- num               //发送数据到channel中（生产）
		fmt.Println("wirte:", idx, num)
		cond.Signal()           //发送信号唤醒阻塞在条件变量上的消费者
		cond.L.Unlock()         //生产结束，解锁互斥锁
		time.Sleep(time.Second) //休息一会
	}

}

//消费者
func consumer(c <-chan int, idx int) {
	//
	for {
		cond.L.Lock()     //给条件变量的互斥锁成员L进行加锁（和生产者是同一个锁）
		for len(c) == 0 { //判断市场空，等待生产者生产
			cond.Wait() //挂起当前goroutine,等待条件满足,被生产者唤醒（解锁并阻塞当前线程,在唤醒返回前加锁）
		}
		fmt.Println("read:", idx, <-c) //接收channle中的数据（消费）
		cond.Signal()                  //发送信号唤醒阻塞在条件变量上的生产者
		cond.L.Unlock()                //消费结束，解锁互斥锁
		time.Sleep(time.Second * 2)    //休息一会
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())

	//初始化条件变量的互斥锁成员L
	cond.L = new(sync.Mutex)

	//创建市场channel（公共缓冲区）
	ch := make(chan int, 5)

	for i := 0; i < 5; i++ { //5个生产者
		go producer(ch, i+1)
	}

	for i := 0; i < 5; i++ { //5个消费者
		go consumer(ch, i+1)
	}

	//主goroutine休息一会
	time.Sleep(time.Second * 5)

}
