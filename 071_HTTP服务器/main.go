package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
	HTTP服务器
*/

func myOpenfile(fileName string, w http.ResponseWriter) {

	//拼接服务器路径和客户端请求的文件名
	filePath := "/Users/waaa/wss/80" + fileName

	if fileName == "/" {
		filePath += "index.html"
	}

	//打开文件
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("<b style=\"color:#f00\">404 Error</b>")) //向客户端发送404
		return
	}
	defer f.Close()

	//读取文件内容
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil && err == io.EOF {
			break
		}
		//向客户端发送文件内容
		w.Write(buf[:n])
	}

}

//回调函数（业务逻辑处理程序），w: 写数据给客户端(响应包)；r: 从客户端读数据(请求包)
func myHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello,haha!!!")) //写入响应包体

	//获取客户端请求的URL（转换为字符转并保存到文件名变量）
	fileName := r.URL.String()

	//调用文件处理函数（传入请求文件名和响应发送接口）
	myOpenfile(fileName, w)

}

func main() {

	//注册处理函数（参数1为服务器提供的URL，参数2为回调函数）
	http.HandleFunc("/", myHandler)

	//服务处理：指定监听地址并调用服务端程序处理连接请求（参数2为nil时，服务端调用 http.DefaultServeMux 进行处理）
	http.ListenAndServe("127.0.0.1:8000", nil)

}
