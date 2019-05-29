package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

/*
	TCP-server（并发服务器）
*/

//4: 封装与客户端交互的函数，参数类型为：net.Conn
func HandlerConnect(conn net.Conn) {

	//8: 延迟关闭连接

	defer conn.Close()

	//获取当前客户端的网络地址信息
	clientAddr := conn.RemoteAddr().String()
	fmt.Println("客户端已连接:", clientAddr)

	//5: 接收客户端的数据

	//创建缓冲区
	buf := make([]byte, 1024)

	for { //循环读取

		n, err := conn.Read(buf)

		if n == 0 { //conn.Read()读到0则表示客户端已关闭
			fmt.Println("客户端已关闭:", clientAddr)
			return //结束当前goroutine
		}
		if err != nil && err != io.EOF {
			fmt.Println("Read:", err) //捕获读取错误
			return
		}

		//设置客户端退出命令
		if string(buf[:n]) == "exit\n" {
			fmt.Println("客户端主动退出:", clientAddr)
			return
		}

		//6: 处理客户端数据（例：小写->大写）

		fmt.Println("client:", clientAddr, strings.ToUpper(string(buf[:n])))

		//7: 发送数据到客户端

		conn.Write([]byte("success\n"))

	}

}

func main() {

	//1: 创建监听器(协议,服务器地址)

	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	//2: 阻塞等待客户端连接请求

	for { //循环监听等待多个客户端
		conn, err := listener.Accept() //返回用于和客户端进行数据通信的socket
		if err != nil {
			fmt.Println(err)
			continue //如果连接出错则跳过本次监听
		}

		//3: 调用客户端数据处理函数，为每一个客户端新建goroutine

		go HandlerConnect(conn)
	}

}
