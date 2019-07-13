/*
	哈希值特性：
		不可逆：无法通过哈希值反推原文内容
		抗碰撞：基数极大，无法通过同一个哈希找到另一个不同的内容
		唯一性：内容不变，哈希值不变，内容改变，哈希值一定改变

	算法示例：
		linux: 	$ md5 [文件名]
				$ sha256 [文件名]
		OSX:	$ shasum -a 256 [文件名]
*/

package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

//Md5hash MD5算法生成哈希值(多次处理内容并计算)
func Md5hash(src []byte) ([]byte, error) {

	//创建一个使用MD5校验的hash.Hash接口
	//接口里实现了io.writer接口
	myHash := md5.New()

	//向接口里写入需要处理的内容
	_, err := io.WriteString(myHash, string(src))
	if err != nil {
		return nil, err
	}

	//计算哈希值
	//Sum接口将接收的参数拼接到生成的哈希值前面，因此这里传入空
	hash := myHash.Sum(nil)

	//返回哈希值
	return hash, nil
}

func main() {
	//原文
	src := []byte("hello~ 哈哈。")

	//计算哈希值
	hash, _ := Md5hash(src)
	fmt.Printf("%x\n", hash) //以16进制打印哈希值

	//一次性计算哈希值（返回值为数组）
	fmt.Printf("%x\n", md5.Sum(src))
}
