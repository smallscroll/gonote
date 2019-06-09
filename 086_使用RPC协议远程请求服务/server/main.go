/*
	RPC（Remote Procedure Call Protocol）—远程过程调用协议，是一种通过网络从远程计算机程序上请求服务，而不需要了解底层网络技术的协议。
	和远程访问或web请求差不多，都是一个client向远端服务器请求服务返回结果，
	但web请求使用的网络协议是http高层协议，而rpc用的协议多为TCP网络层协议，减少了信息的包装，加快了处理速度。
*/

package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

/*
	RPC服务
*/

//Calc Calc服务类
type Calc struct {
}

//Add Calc服务的方法
func (c *Calc) Add(In []float64, Out *float64) error {
	var result float64
	for _, v := range In {
		result += v
	}
	*Out = result
	return nil
}

func main() {
	//实例化服务对象
	c := new(Calc)

	//注册rpc服务
	rpc.Register(c)

	//启动rpc网络
	rpc.HandleHTTP()

	//创建监听器
	listener, err := net.Listen("tcp", "127.0.0.1:8086")
	if err != nil {
		fmt.Println(err)
		return
	}

	//服务处理
	http.Serve(listener, nil)

}
