package main

import (
	"fmt"
	"net"
)

/*
	创建http客户端接收并打印http相应包
*/

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	//伪装浏览器发送http请求报文(\r\n分隔与结束)
	httpRequest := "GET /hello HTTP/1.1\r\nHost:127.0.0.1:8000\r\n\r\n"
	//发送请求包
	conn.Write([]byte(httpRequest))

	//获取响应包
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf[:n])) //打印响应包内容
}

/*
	HTTP响应报文格式：
*/

//1. 响应行（协议版本 状态码 状态描述）

// HTTP/1.1 200 OK

//2. 响应头（key:value）

// Date: Sun, 07 Apr 2019 04:18:37 GMT
// Content-Length: 13
// Content-Type: text/plain; charset=utf-8

//3. 空行（\r\n）表示http响应头结束
//

//4. 响应包体（服务器返回的数据或错误信息）
// Hello,haha!!!
