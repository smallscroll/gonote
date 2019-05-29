package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// //设定文件权限掩码
	// retMask := syscall.Umask(0)
	// fmt.Println(retMask)

	//创建文件file1，如果文件已存在，自动截断（清空），返回值为文件指针
	f, err := os.Create("file1")
	if err != nil {
		fmt.Println(err)
		return
	}
	//关闭文件
	defer f.Close()

	//按字符串写文件（返回写入内容字节数和错误信息）
	n, err := f.WriteString("Hello,world.\n")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d success\n", n)

	// //以只读方式打开文件
	// f, err = os.Open("file1")

	//以读写方式打开文件(文件,读写模式,参数三：打开权限(当os.O_CREATE时必须指定参数0-7(0666/0777...)，否则都可以指定0))
	f, err = os.OpenFile("file1", os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	//以读写和追加模式打开文件(如果文件不存在则创建文件)
	f, err = os.OpenFile("file1", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	//按字符串写文件，返回写入内容字节数和错误信息（多次写入光标自动后移）
	n, err = f.WriteString("hahah~\n")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d success\n", n)

	/*
		按位置写入文件
	*/

	//按位置写文件
	//修改文件读写指针位置(偏移量,偏移开始开始位置)
	//返回从文件起始位置到当前文件读写指针位置的偏移量和错误信息
	ret, _ := f.Seek(2, io.SeekStart) //文件读写指针从文件起始开始，偏移2

	//在偏移后的位置写入(写入数据(字符切片),偏移量)
	//写入内容会根据写入长度覆盖后面对应位置的数据
	n, err = f.WriteAt([]byte("nihao"), ret)
	if err != nil {
		fmt.Println("WriteString err", err)
		return
	}
	fmt.Printf("%d success\n", n)

	//文件读写指针从文件末尾开始，偏移10
	ret, _ = f.Seek(10, io.SeekEnd)
	n, err = f.WriteAt([]byte("hehe\n"), ret)
	if err != nil {
		fmt.Println("WriteString err", err)
		return
	}
	fmt.Printf("%d success\n", n)

	f.Close()

	/*
	   按行读文件
	*/

	//只读打开文件
	f, _ = os.Open("file1")

	//创建阅读器reader（自带缓冲区），把文件内容放入缓冲区
	reader := bufio.NewReader(f)

	//创建用来接收已读取数据的buf
	buf := make([]byte, 1024)

	for {
		//从reader缓冲区中读取数据，以\n作为读取结束标记
		buf, err = reader.ReadBytes('\n')
		//当读到文件结尾时，err会被置为EOF，表示文件读取结束
		if err != nil && err == io.EOF {
			fmt.Println("fsinish")
			break
		}
		fmt.Print(string(buf))
	}

	// //删除文件
	// os.Remove("file1")
}
