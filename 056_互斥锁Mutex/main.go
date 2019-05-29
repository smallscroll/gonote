package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	互斥锁：建议锁，不具有强制性，保护公共区被锁住后只有成功加锁的go程访问，其他go程阻塞在锁的等待上
	加锁：公共区访问开始之前
	解锁：访问结束后立即解锁
*/

//定义互斥锁
var mutex sync.Mutex

func printer(str string) {
	mutex.Lock() //加锁
	for _, word := range str {
		fmt.Printf("%c", word)
		time.Sleep(time.Second / 5)
	}
	mutex.Unlock() //解锁

	//...
}

func user1() {
	printer("hello")
}

func user2() {
	printer("world")
}

func main() {

	go user1()
	go user2()

	time.Sleep(time.Second * 3)
	fmt.Println("")

}
