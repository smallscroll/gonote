package main

import (
	"fmt"
	"net"
)

/*
	创建http服务器接收并打印http请求包
*/

func main() {
	//创建监听器
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	//等待客户端连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	//接收请求包
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if n == 0 {
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v", string(buf[:n])) //打印请求包内容
}

/*
	HTTP请求报文格式
*/

//1. 请求行（请求方法 URL 协议版本）

// GET / HTTP/1.1

//2. 请求头（key: value）

// Host: 127.0.0.1:8080
// Upgrade-Insecure-Requests: 1
// Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
// User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.1 Safari/605.1.15
// Accept-Language: zh-cn
// Accept-Encoding: gzip, deflate
// Connection: keep-alive

//3. 空行（\r\n ）表示http请求头结束

//

//4. 请求包体（GET方法没有包体，POST方法有包体）
