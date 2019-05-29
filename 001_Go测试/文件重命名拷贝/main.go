package main

/*
	将指定a/目录下的所有MP3文件提取，重命名为1xx.mp3，2xx.mp3，3xx.mp3 ..... 拷贝到b/目录下。
*/

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//定义目录全局变量
var oldPath = "/Users/waaa/Downloads/a/"
var newPath = "/Users/waaa/Downloads/b/"

//定义文件拷贝函数
func copyFile(oldFile string) {
	//只读打开旧文件
	of, err := os.Open(oldPath + oldFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer of.Close()

	//创建新文件
	nf, err := os.Create(newPath + oldFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer nf.Close()

	//创建数据缓冲区
	buf := make([]byte, 4096)

	//循环读取旧文件
	for {
		n, err := of.Read(buf)
		if n == 0 {
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		}
		//将读取数据写入新文件
		nf.Write(buf[:n])
	}

}

func main() {

	//只读打开旧目录
	of, err := os.OpenFile(oldPath, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer of.Close()

	//读取目录，获得目录项
	infos, err := of.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}
	//定义文件名计数
	var i int
	//遍历目录项
	for _, info := range infos {
		//筛选出普通文件
		if info.IsDir() == false {
			//筛选mp3文件
			if strings.HasSuffix(info.Name(), ".mp3") {
				//重命名文件
				os.Rename(oldPath+info.Name(), oldPath+strconv.Itoa(i+1)+"xx.mp3")
				//调用文件拷贝函数(只传入文件名)
				copyFile(strconv.Itoa(i+1) + "xx.mp3")
				i++
			}
		}
	}
}
