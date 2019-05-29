package main

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

/*
	序列化与反序列化
*/

type student struct {
	Name string
	Age  int
}

func main() {
	//初始化一个结构体切片
	stus := []student{{"jack", 18}, {"yoyo", 20}}

	//连接数据库
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("连接失败", err)
	}
	defer conn.Close()

	/*
		序列化（字节化）: 编码
	*/

	//容器：定义一个容器接收二进制数据
	var buffer bytes.Buffer
	//编码器：传入容器地址，返回一个编码器
	enc := gob.NewEncoder(&buffer)
	//编码：把结构体切片对象编码为二进制数据
	enc.Encode(stus)

	//存入数据库
	conn.Do("set", "stus", buffer.Bytes())

	/*
		反序列化（反字节化）: 解码
	*/

	resp, _ := conn.Do("get", "stus")
	result, err := redis.Bytes(resp, err) //将返回结果转化为字节切片
	//解码器：
	dec := gob.NewDecoder(bytes.NewReader(result))
	//容器：定义一个容器接收解码后的数据
	var newStus []student
	//解码：将解码后的数据存到结结构体切片对象的容器(结构体成员首字母需大写)
	dec.Decode(&newStus)

	fmt.Println(newStus)
}
