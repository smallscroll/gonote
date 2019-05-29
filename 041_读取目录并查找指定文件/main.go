package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	/*
	   目录操作
	*/

	//读目录并筛选jpg文件

	fmt.Println("输入要查找的目录：")
	var path string
	fmt.Scan(&path)

	//打开目录，参数二选择只读，参数三必须选择目录模式，返回值为目录类型文件指针
	dir, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dir.Close()

	//读取目录，传递待读取项目个数，返回值为目录项切片（-1为读取全部目录项）
	infos, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}

	////范围遍历目录项切片
	for _, infos := range infos {
		//筛选非目录文件，返回值为bool
		if !infos.IsDir() {
			////筛选后缀名为.jpg的文件，返回布尔值
			if strings.HasSuffix(infos.Name(), ".jpg") {
				//打印文件名
				fmt.Println(infos.Name())
			}

		}
	}

	/*
		创建目录
	*/

	//使用指定的权限和名称创建一个目录，如果目录已存在返回nil
	os.MkdirAll("./testdir", 0777)

	//使用指定的权限和名称创建一个目录，如果目录已存在会报错。
	err = os.Mkdir("./testdir", 0777)
	if err != nil {
		fmt.Println(err)
	}

}
