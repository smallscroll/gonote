package main

import (
	"fmt"
	"io"
	"net/http"
)

/*
	HTTP客户端
*/

func main() {
	//调用Get()，传入请求的URL，获取响应包(结构体)
	response, err := http.Get("http://127.0.0.1:8000/hello")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close() //关闭响应包体

	// fmt.Println(response.Status) //响应包状态

	//循环读取响应包体

	buf := make([]byte, 4096)

	var result string //定义保存内容的字符串

	for {
		n, err := response.Body.Read(buf)
		if n == 0 { //读取完成
			break
		}
		if err != nil && err != io.EOF { //读取错误
			// panic(err)
			fmt.Println(err)
			return
		}
		fmt.Println(string(buf[:n])) //打印响应包体内容

		result += string(buf[:n]) //将内容保存到字符串
		fmt.Println(result)       //打印字符串
	}
}
