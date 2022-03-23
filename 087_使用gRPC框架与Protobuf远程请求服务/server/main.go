/*

	gRPC: 一个高性能、开源和通用的RPC框架，面向移动和 HTTP/2 设计。

	使用grpc前须安装的依赖包：
		git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc
		git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto
		git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
		git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text


	编译proto文件时须指定grpc插件进行编译：
    	示例：$ protoc --go_out=plugins=grpc:. *.proto

*/

package main

import (
	"context"
	"fmt"
	pb "gonote/087_使用gRPC框架与Protobuf远程请求服务/myproto"
	"net"

	"google.golang.org/grpc"
)

//Calc 服务对象
type Calc struct {
	In  float64
	Out float64
}

//Add 实现RPC中Calc服务接口的方法
func (c *Calc) Add(ctx context.Context, In *pb.In) (*pb.Out, error) {
	var result float64
	for _, v := range In.Nums {
		result += v
	}
	return &pb.Out{Result: result}, nil
}

func main() {
	//创建监听器
	listener, err := net.Listen("tcp", "127.0.0.1:8086")
	if err != nil {
		fmt.Println(err)
		return
	}

	//创建grpc服务
	server := grpc.NewServer()

	//将Calc注册到grpc服务中
	pb.RegisterCalcServer(server, &Calc{})

	//服务处理
	err = server.Serve(listener)
	if err != nil {
		fmt.Println(err)
		return
	}
}
