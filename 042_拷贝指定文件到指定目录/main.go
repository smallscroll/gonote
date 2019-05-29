package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

/*
	拷贝所有mp3文件到指定目录
*/

var path = "./"         //原目录
var ndir = "./testdir/" //新目录

func copymp3(mp3name string, ndir string) {
	//只读方式打开文件（须包含路径的完整文件）
	f, err := os.Open(path + mp3name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	//创建新目录，如果目录已存在返回nil
	os.MkdirAll(ndir, 0777)

	//在新目录下创建新文件
	f2, err := os.Create(ndir + mp3name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f2.Close()

	//创建缓冲区来存储读物文件的内容
	buf := make([]byte, 1024)

	//循环读取文件
	for {
		//按缓冲区大小循环读取文件内容并放入缓冲区，但会读取内容的字节长度
		n, err := f.Read(buf)
		if err != nil && err == io.EOF { //判断读取到文件末尾后跳出
			fmt.Println(err)
			break
		}
		//将读取到的内容按字节写入到新文件
		f2.Write(buf[:n])

	}

}

func main() {
	//打开原目录
	odir, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer odir.Close()

	//读取目录，返回值为目录项且切片（-1为读取全部目录项）
	infos, err := odir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}

	//范围遍历目录项切片
	for _, info := range infos {
		//筛选非目录文件
		if !info.IsDir() {
			//筛选后缀名为.mp3的文件
			if strings.HasSuffix(info.Name(), ".mp3") {
				//将文件名和新目录传入拷贝函数
				copymp3(info.Name(), ndir)
			}
		}
	}
	fmt.Println("finish")
}
