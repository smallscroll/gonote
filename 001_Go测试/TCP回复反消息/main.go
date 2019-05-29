package main

import (
	"fmt"
	"net"
	"strings"
)

/*
	编一个TCP并发服务器，可同时支持N个客户端访问。服务器接收客户端发送内容，将内容按单词逆置，回发给客户端。 如： 客户端发送：this is a socket test 服务器回复：test socket a is this
*/

//定义通信函数

func connControl(conn net.Conn) {
	defer conn.Close()

	//循环读取客户端发送的数据
	for {
		//创建数据缓冲区
		buf := make([]byte, 1024)
		//读取数据
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("客户端关闭")
			return
		}
		if err != nil {
			fmt.Println(err)
			return
		}

		//处理客户端数据
		str := string(buf[:n])
		//按空格切割字符串
		slice := strings.Fields(str)
		//字符串切片逆置
		sliceStart := 0
		sliceEnd := len(slice) - 1
		for {
			if sliceStart >= sliceEnd {
				break
			}
			slice[sliceStart], slice[sliceEnd] = slice[sliceEnd], slice[sliceStart]
			sliceStart++
			sliceEnd--
		}
		//循环发送字符串数据给客户端
		for _, v := range slice {
			_, err := conn.Write([]byte(v + " "))
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		//字符串发送完后再给客户端发送个换行符
		_, err = conn.Write([]byte("\n"))
		if err != nil {
			fmt.Println(err)
			return
		}

	}

}

func main() {

	//创建监听器
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	//循环阻塞等待客户端连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		//调用通信函数
		go connControl(conn)
	}

}
