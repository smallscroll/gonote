package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
	统计指定目录下所有txt文件中所有love单词的个数
*/

var path = "./" //目录

func findlove(file string) int {
	//打开接收到的文件
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer f.Close()

	//创建阅读器缓冲区用于接收文件数据
	reader := bufio.NewReader(f)

	//定义文件内单词个数
	var nums int = 0

	//循环读取文件内容
	for {
		//读取文件内容，以换行分隔符作为每次读取结束标志
		str, err := reader.ReadBytes('\n')
		if err != nil && err == io.EOF { //判断读取到文件末尾时结束循环
			break
		}
		//将每次读取的内容传入函数lovesnum进行分割和单词统计,接收love个数进行累加
		nums += lovesnum(string(str))
	}
	//返回文件内单词个数
	return nums
}

func lovesnum(str string) int {
	//将字符串按空格符进行切割，得到单词切片
	words := strings.Fields(str)
	//创建以字符串作为key的map
	m := make(map[string]int)
	//循环遍历单词切片，将键为love的值进行累加，从而得到传入内容里love的个数
	for i := 0; i < len(words); i++ {
		m[words[i]]++
	}
	//返回结果
	return m["love"]
}

func main() {
	//打开指定目录
	dir, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("OpenFile", err)
		return
	}
	defer dir.Close()

	//读取目录，返回值为目录项切片
	infos, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Readdir", err)
		return
	}

	//初始化单词总个数
	var loves int = 0

	//遍历目录项
	for _, info := range infos {
		//筛选非目录文件
		if !info.IsDir() {
			//筛选后缀名为.txt的文件
			if strings.HasSuffix(info.Name(), ".txt") {
				//将文件路径传入findlove函数，并接收单词个数进行；累加
				loves += findlove(path + info.Name())
			}
		}
	}

	fmt.Println("love:", loves)
}
