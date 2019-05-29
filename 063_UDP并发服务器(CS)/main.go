package main

import (
	"fmt"
	"net"
)

/*

	TCP 与 UDP

	TCP是面向连接的可靠的数据包传递；针对不稳定的网络层做完全弥补（丢包后借助回执重传）
		• 稳定（有序、速度）、可靠
		• 系统资源占用多（需要维护连接）、发送速度慢、开发难度大

	用于对数据稳定性、准确性要求较高的场合，例如：上传、下载


	UDP是无连接的不可靠报文传输；针对不稳定的网络层直接还原真实状态（丢包后不处理）
		• 系统资源占用小（需要维护连接）、发送速度快、开发难度小
		• 不稳定（无序、慢速）、不可靠

	用于对数据的传输速度要求较高的场合，允许适当的数据丢失，例如：游戏、直播、电话会议、视频


	详细对比：

	TCP										UDP
	面向连接						  			面向无连接
	要求系统资源较多							要求系统资源较少
	TCP程序结构较复杂							UDP程序结构较简单
	使用流式									使用数据包式
	保证数据准确性								不保证数据准确性
	保证数据顺序								不保证数据顺序
	通讯速度较慢								通讯速度较快

*/

/*
	UDP-server
*/

func main() {

	//1: 指定udp协议创建监听的地址（服务器UDP地址结构）
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println(err)
		return
	}

	//2: 等待客户端连接，创建数据通信Socket
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer udpConn.Close()

	//3: 接收客户端发送的数据

	//创建缓冲区
	buf := make([]byte, 4096)

	for { //循环读取数据

		//UDP服务器并不需要并发，通过循环处理客户端数据即可

		n, clientAddr, err := udpConn.ReadFromUDP(buf) //返回客户端地址
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("client:", clientAddr, string(buf[:n]))

		go func() {
			//4: 发送数据给客户端
			_, err = udpConn.WriteToUDP([]byte("success\n"), clientAddr) //指定客户端地址
			if err != nil {
				fmt.Println(err)
				return
			}
		}()

	}

	//5: 关闭(defer)

}

/*
	//linux模拟客户端测试
	$nc -u 127.0.0.1 8001

	...
*/
