package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	读写锁：读共享，写独占；写锁优先级高于读锁
	锁只有一把，但是具备两种加锁属性(读/写)
*/

//创建全局变量作为公共区
var count int //默认值为0

//定义读写锁
var rwmutex sync.RWMutex

var ch chan int

func readGo(idx int) {
	for {
		//读模式加锁
		rwmutex.RLock()
		fmt.Println("read:", idx, count)
		//读模式解锁
		rwmutex.RUnlock()
		time.Sleep(time.Second * 2)
	}
}

func writeGo(idx int) {
	for {
		//写模式加锁
		rwmutex.Lock()
		count = rand.Intn(500)
		fmt.Println("write", idx, count)
		//写模式解锁
		rwmutex.Unlock()
		time.Sleep(time.Second)
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		go readGo(i + 1)
	}

	for i := 0; i < 5; i++ {
		go writeGo(i + 1)
	}

	time.Sleep(time.Second * 5)
	fmt.Println("")

}
