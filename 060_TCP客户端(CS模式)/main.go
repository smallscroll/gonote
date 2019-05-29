package main

import (
	"fmt"
	"net"
)

/*
	TCP-client
*/

func main() {

	//1: 连接服务器(使用协议,服务器地址(IP:PORT))

	conn, err := net.Dial("tcp", "127.0.0.1:8000") //返回用于和服务器进行数据通信的socket //Dial()函数返回表示连接成功
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close() //延迟关闭连接

	//2: 发送(写)数据给服务器

	for { //循环发送
		var tmp string
		fmt.Scan(&tmp)
		n, err := conn.Write([]byte(tmp))
		if err != nil {
			fmt.Println(err)
			return
		}

		//3: 接收(读)服务器回发的数据

		//创建缓冲区
		buf := make([]byte, 1024)

		//读取数据
		n, err = conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("server:", string(buf[:n]))
	}

	//4: 关闭(defer)

}
