package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	// //GOMAXPROCS()设置可同时执行的最大CPU数，并返回先前的设置
	// runtime.GOMAXPROCS(2)

	/*
		当创建goroutine后，main()便作为主goroutine存在
		主goroutine退出后会释放虚拟内存地址空间，因此其它的子goroutine也会被结束
		注：当gotoutine的主函数（非main函数）提前退出时，如果goroutine还没结束，其主函数的主函数会接管该goroutine，直到goroutine自己结束或者最后一个接管者（main函数）退出时将goroutine结束
	*/

	//创建goroutine
	go func(str string) {
		for i := 0; i < 3; i++ {
			defer fmt.Println("over...")
			fmt.Println(str)
			time.Sleep(time.Second)
			// //Goexit()将终止调用它的goroutine，Goexit()会在终止该goroutine前执行已注册的defer函数
			// //不能在主goroutine中使用Goexit()来结束主goroutine，否则会出现死锁异常
			runtime.Goexit()
		}
	}("ding..ding...")

	for i := 0; i < 3; i++ {
		fmt.Println("main..main...mian....")
		time.Sleep(time.Second)
		//Gosched()用于让出CPU时间片，让出当前主goroutine的执行权限，调度器安排其他等待的任务运行，
		//并在下次再获得cpu时间轮片的时候，从该出让cpu的位置恢复执行。
		runtime.Gosched()
	}

}
