package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

/*
	TCP-client（并发客户端）
*/

func main() {

	//1: 连接服务器(使用协议,服务器地址(IP:PORT))

	conn, err := net.Dial("tcp", "127.0.0.1:8000") //返回用于和服务器进行数据通信的socket
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

/*

TCP通信过程：

	三次握手：
		主动建立连接请求端，发送SYN标志位(携带序号)
		被动接受连接请求端，接收SYN标志位，回发ACK(携带确认序号)；同时发送SYN标志位(携带序号)
		主动建立连接请求端，接收SYN标志位，回发ACK(携带确认序号)；
			---3次握手完成对应的应用层：
				客户端：Dial()返回
				服务器：Accept()返回

TCP数据通信：

	发送端发送数据包(携带序号)
	接收端收到数据包，给发送端发送ACK应答(确认序号)
	接收端采用批量回执、发送滑动窗口(实时告知本端储存数据的缓冲区大小)
	...
		---如果丢失数据包或ACK，TCP协议将自动重发缓冲区的数据包


四次挥手：

	主动断开连接请求，发送FIN标志位(携带序号)
	被动断开连接请求，接收FIN标志位，回发ACK应答(携带确认序号)
		---半关闭完成
	被动断开连接请求，发送FIN标志位(携带序号)
	主动断开连接请求，接收FIN标志位，回发ACK应答(携带确认序号)
		---最后一个ACK被接收后，4次挥手完成关闭


TCP状态转换：

	主动连接端：
		CLOSED --> 发送SYN --> SYN_SENT --> 接收ACK、SYN，发送ACK --> ESTABLISHED --> 数据通信
	主动关闭端：
		ESTABLISHED --> 发送FIN --> FIN_WAT_1 --> 接收ACK --> FIN_WAIT_2（半关闭） --> 接收FIN，发送ACK --> TIME_WAIR --> 等待2MSL时长 --> CLOSED

	(2MSL：为了确保TCP通信过程中的最后一个ACK能被对端收到，等待一个时长（约40s）)


	被动连接端：
		CLOSED --> LISTEN，接收SYN，发送ACK、SYN --> SYN_RCVD -->接收ACK --> ESTABLISHED --> 数据通信
	被动关闭端：
		ESTABLISHED --> 接收FIN，发送ACK --> CLOSE_WAIT（对应主动端的半关闭）--> 发送FIN --> LAST_ACK --> 接收ACK --> CLOSED

*/
