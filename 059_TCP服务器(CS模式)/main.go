package main

import (
	"fmt"
	"net"
)

/*

	网络分层架构：

	OSI/RM七层模型（理论标准）：
		•••• 	应用层
		•••• 	表示层
		•••• 	会话层
		••• 	传输层
		••		网络层
		• 		数据链路层
		• 		物理层


	TCP/IP四层模型（实际标准）：

		应用层：FTP; Telnet; TFTP; NFS; HTTP
		传输程：TCP(传输控制协议); UDP(用户数据报协议)
		网络层：IP(internet协议); ICMP(internet控制报文协议); IGMP(internet组管理协议)
		链路层：ARP(正向地址解析协议); RARP(反向地址解析协议)

		(传输层：TCP类似打电话，UDP类似发短信)


	Socket：套接字
		伪文件，用于描述IP地址和端口，可以实现不同程序间的数据通信
		Socket内部相当于两个双向channel(分别用于读/写)
		在一次网络通信过程中，参与通信的socket必须“成对”出现


	网络应用程序设计模式：
		C/S：Client/Server 优点：缓存大量数据，协议选择灵活；缺点：用户安全性差，跨平台能力差，开发工作量大
		B/S：Browser/Server ...

*/

/*
	TCP-server
*/

func main() {

	//1: 创建监听器(使用协议,服务器地址(IP:PORT))

	listener, err := net.Listen("tcp", "127.0.0.1:8000") //接口类型
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	//2: 阻塞等待客户端连接

	conn, err := listener.Accept() //返回用于和客户端进行数据通信的socket	//接口类型	//Accept()函数返回表示连接成功
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close() //延迟关闭连接

	//3: 接收(读)客户端的数据

	//创建缓冲区
	buf := make([]byte, 1024)

	for { //循环读取
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("client:", string(buf[:n]))

		//4: 发送(写)数据到客户端

		conn.Write([]byte("success\n"))

	}

	//5: 关闭(defer)

	//
	/*
		//linux模拟客户端测试（netcat）
		nc 127.0.0.1 8000

		//linux查看进程
		netstat -anp | grep 8000
		//mac查看进程
		lsof -i :8000

		...
	*/

}
