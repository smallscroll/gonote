package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

/*
	网络文件传输：接收端

	1.创建监听
	2.创建通信Socket
	3.接收客户端发送的文件名并保存
	4.回发ok给客户端
	5.封装调用文件接收函数
		创建文件-读取数据-写入文件-判断结束并关闭

*/

func main() {

	//1: 创建监听器
	listener, err := net.Listen("tcp", "127.0.0.1:8002")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close() //延迟关闭

	//2: 与客户端建立通信

	//创建通信socket
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close() //延迟关闭

	//接收文件名
	buf := make([]byte, 4096)

	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileName := string(buf[:n])

	//回发信息确认通信成功
	conn.Write([]byte("ok"))

	//3: 调用文件接收函数
	recvFile(fileName, conn)

}

func recvFile(fileName string, conn net.Conn) {

	//4: 创建新文件并接收文件数据

	path := "/users/waaa/Downloads/" //设置路径

	//创建文件
	f, err := os.Create(path + fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close() //延迟关闭

	//创建缓冲区
	buf := make([]byte, 4096)

	//循环接收数据
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("拷贝文件完成")
			return
		}
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		}

		//	写入数据到新文件
		f.Write(buf[:n])
	}

}
