package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/*
	   大文件拷贝（按字节（块）读取）
	*/

	//只读打开原文件

	fo, err := os.Open("file2")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fo.Close()

	//创建新文件

	fn, err := os.Create("file3")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fn.Close()

	//创建存储数据的临时切片缓冲区
	buf := make([]byte, 1024)

	//循环读文件数据到缓冲区
	for {
		//调用read()方法，按缓冲区字节大小读取原文件放入临时缓冲区，返回读取内容的字节长度
		n, err := fo.Read(buf)
		if n == 0 {
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		}
		//将临时缓冲区的内容按字节写入到新文件（读了多少就写入多少），返回写入内容的字节长度
		_, err = fn.Write(buf[:n])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n)
	}
	fmt.Println("Copy finish")

}
