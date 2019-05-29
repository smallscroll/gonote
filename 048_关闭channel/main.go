package main

import (
	"fmt"
)

func main() {

	//创建缓冲区大小为3的channel
	ch := make(chan int, 3)

	go func() { //子goroutine
		for i := 0; i < 7; i++ {
			ch <- i
			fmt.Println("...", len(ch), cap(ch))
		}
		//不需要再发送数据时，关闭channel
		close(ch)
		//已关闭的channel将不能再发送数据，但可以从它那继续接收数据
		//从已关闭的channel接收到的数据为数据类型的默认值
	}()

	//主goroutine

	////判断channel是否关闭（方式1）
	// for {
	// 	//接收通道数据，赋值给变量并判断变量真假（true表示通道未关闭）
	// 	if num, ok := <-ch; ok {
	// 		fmt.Println(num)
	// 	} else { //如果ok为false表示没有接收到数据：通道关闭
	// 		fmt.Println(ok)
	// 		break
	// 	}
	// }

	//判断channel是否关闭（方式2）
	for v := range ch { //range一个通道类型等同于接收通道数据
		fmt.Println("---", v)
	}

}
