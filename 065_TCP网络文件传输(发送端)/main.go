package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

/*
	网络文件传输：发送端

	1.创建客户端的通信的Socket
	2.通过命令行参数得到文件路径
	3.获取文件属性，得到文件名
	4.发送文件名到接收端
	5.读取并判断接收端回发的ok
	6.封装并调用文件发送函数
		打开文件-循环读取文件-发送文件到接收端-判断结束-关闭

*/

func main() {

	//1: 获取文件名

	list := os.Args //获取命令行参数（索引0的值为程序本身）

	//判断至少输入一个参数
	if len(list) < 2 {
		fmt.Println("请输入文件名(格式：go run xxx.go 文件名(包含路径))")
		return
	}

	//获取文件路径
	filePath := list[1]

	//获取文件属性
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println(err)
	}
	//获取文件名
	fileName := fileInfo.Name()

	//2: 连接服务器

	//创建通信Socket
	conn, err := net.Dial("tcp", "127.0.0.1:8002")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close() //延迟关闭

	//3: 与服务器确认通信成功

	//发送文件名给服务器
	conn.Write([]byte(fileName))

	//读取服务器返回数据
	buf := make([]byte, 4069)

	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	if string(buf[:n]) == "ok" {
		//3: 调用文件发送函数
		sendFile(filePath, conn)
	} else {
		fmt.Println("通信失败")
		return
	}

}

//封装文件发送函数
func sendFile(filePath string, conn net.Conn) {

	//4: 读取并发送文件

	//打开文件
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close() //延迟关闭

	//创建缓冲区
	buf := make([]byte, 4096)

	//循环读取
	for {
		n, err := f.Read(buf)
		if n == 0 {
			fmt.Println("发送文件完成")
			return
		}
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		}

		//发送数据到服务器
		conn.Write(buf[:n])
	}

}
