package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

/*
	UDP-client
*/

func main() {

	//1: 连接服务器(使用协议,服务器地址(IP:PORT))

	conn, err := net.Dial("udp", "127.0.0.1:8001") //返回用于和服务器进行数据通信的socket
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close() //延迟关闭连接

	//2: 发送数据给服务器

	//创建缓冲区用于储存用户键盘输入的数据
	str := make([]byte, 1024)

	go func() { //创建子goroutine用于读取用户数据并发送

		for { //循环读取并发送

			//读取用户键盘输入
			n, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println(err)
				return
			}

			//发送数据到服务器
			n, err = conn.Write([]byte(str[:n]))
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}()

	//3: 接收服务器回发的数据（主goroutine）

	//创建缓冲区用于储存服务器回发的数据
	buf := make([]byte, 1024)

	for {

		//读取数据
		n, err := conn.Read(buf)
		if n == 0 { //接收到0表示服务器端关闭
			fmt.Println("连接已断开")
			return
		}
		if err != nil && err != io.EOF { //捕获读取错误
			fmt.Println(err)
			return
		}
		fmt.Println("server:", string(buf[:n]))
	}

	//4: 关闭(defer)

}
