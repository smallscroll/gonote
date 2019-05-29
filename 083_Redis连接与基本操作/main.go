package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	//连接数据库
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("连接失败", err)
	}
	defer conn.Close()

	//操作1
	//Send函数发送指令到缓冲区，Flush将缓冲区命令刷新到服务器，Receive接收服务器返回的数据
	conn.Send("set", "key1", "hello,haha~")
	conn.Flush()
	resp1, _ := conn.Receive()
	fmt.Println("操作1:", resp1)

	//操作2
	//Do函数直接执行命令
	resp2, _ := conn.Do("get", "key1")
	//回复助手函数：将回复结果转换为指定类型的值
	result2, _ := redis.String(resp2, err)
	fmt.Println("操作2:", result2)

	/*
		示例3: 获取多个不同类型的数据
	*/
	//设置键为k1值为1000、键为k2值为Hello~
	conn.Do("mset", "k1", "1000", "k2", "Hello~")

	//获取数据
	resp3, _ := conn.Do("mget", "k1", "k2") //获取多个键值
	result3, _ := redis.Values(resp3, err)  //转换回复结果，返回接口切片类型的数据
	fmt.Printf("示例3多值获取: %c\n", resp3)
	//通过Scan函数把数据转换并保存到对应类型的变量里
	var v31 int
	var v32 string
	redis.Scan(result3, &v31, &v32) //如果数据类型和变量类型不匹配，那么变量值会显示该类型变量的默认值
	fmt.Println("示例3变量值:", v31, v32)
	fmt.Printf("示例3变量类型: %T,%T\n", v31, v32)
}
