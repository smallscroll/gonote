package main

import (
	"fmt"
	"net/rpc"
)

func main() {

	//连接rpc服务
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8086")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	//定义接收返回参数的变量
	var result float64

	//远程调用函数(被调用的rpc服务方法，传入的参数，返回的参数)
	err = client.Call("Calc.Add", []float64{5, 7}, &result)
	if err != nil {
		fmt.Println(err)
		return
	}

	//打印结果
	fmt.Println(result)
}
