package main

import (
	"context"
	"fmt"

	pb "test/myproto"

	"google.golang.org/grpc"
)

func main() {
	//连接grpc服务
	conn, err := grpc.Dial("127.0.0.1:8086", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	//创建grpc客户端
	c := pb.NewCalcClient(conn)

	//远程调用Calc接口方法
	out, err := c.Add(context.Background(), &pb.In{Nums: []float64{5, 7, 9}})
	if err != nil {
		fmt.Println(err)
		return
	}

	//打印结果
	fmt.Println("计算结果:", out.Result)

}
