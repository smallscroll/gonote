package main

import (
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
	"time"
)

/*

	网络聊天室-server

*/

//创建客户端结构体
type Client struct {
	name string
	addr string
	C    chan string //客户端消息通道
}

//创建在线客户端列表(客户端地址作为键)
var clientList = make(map[string]Client)

//创建读写锁(保护全局map变量)
var rwMutex sync.RWMutex

//创建全局消息通道
var messageChannle = make(chan string)

//封装消息广播函数
func messageToAll() {
	//循环接收
	for {
		//阻塞等待接收全局消息通道的数据
		msg := <-messageChannle

		//遍历并将消息发送给所有在线客户端

		rwMutex.RLock() //遍历全局map之前加读锁

		for _, client := range clientList {
			client.C <- msg //将数据发送到客户端自带的消息通道
		}

		rwMutex.RUnlock() //遍历完后立即解锁
	}

}

//封装客户端通信函数
func messageToClient(client Client, conn net.Conn) {

	for msg := range client.C { //循环遍历客户端消息通道，如果没有数据自动结束循环
		conn.Write([]byte(msg)) //将数据发送给客户端
	}
}

//封装消息拼接函数
func makeMsg(client Client, msg string) string {
	return "[" + client.addr + "]" + client.name + ":" + msg
}

//封装客户端通信控制函数
func clientControl(conn net.Conn) {

	defer conn.Close()

	//获取客户端地址
	clientAddr := conn.RemoteAddr().String()
	//创建客户端对象(将客户端地址作为name默认值)
	client := Client{clientAddr, clientAddr, make(chan string)}
	//将客户端添加到在线列表
	rwMutex.Lock() //全局map写锁
	clientList[clientAddr] = client
	rwMutex.Unlock() //解锁

	//创建客户端通信goroutine(接收客户端消息通道的信息并发送给客户端)
	go messageToClient(client, conn)

	//广播客户端上线消息（将消息发送到全局消息通道）
	messageChannle <- makeMsg(client, "login\n")

	//创建客户端退出监测通道
	var exitChannel = make(chan bool)
	//创建数据监测通道以重置After定时器
	var hasData = make(chan bool)

	//创建客户端消息处理goroutine
	go func() {
		//创建缓冲区
		buf := make([]byte, 4096)
		//循环接收客户端消息
		for {
			n, err := conn.Read(buf)
			if n == 0 { //检测客户端关闭
				fmt.Println(client.addr + ":exit")
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn.Read", err)
				return
			}
			//处理客户端发送的数据
			msg := string(buf[:n])

			//消息命令处理
			if msg == "who\n" { //查询在线列表
				rwMutex.RLock() //全局map读锁
				for _, client := range clientList {
					online := makeMsg(client, "---online\n")
					conn.Write([]byte(online)) //将列表信息发给客户端
				}
				rwMutex.RUnlock() //解锁
			} else if len(msg) > 7 && msg[:7] == "rename " { //修改用户名命令格式：rename [新用户名]
				newname := strings.Split(msg, " ")[1]         //按空格拆分字符串
				client.name = strings.Split(newname, "\n")[0] //修改当前对象的name值

				rwMutex.Lock()                   //全局map写锁
				clientList[client.addr] = client //将当前对象更新到全局列表
				rwMutex.Unlock()                 //解锁

				conn.Write([]byte("your new name:" + client.name + "\n")) //将修改成功信息发送给客户端
			} else if msg == "exit\n" {
				exitChannel <- true
				return
			} else {
				//将消息发送到全局消息通道进行广播
				messageChannle <- makeMsg(client, msg)

			}

			hasData <- true //保持数据监测通道畅通

		}
	}()

	//监测通道
	for {
		select {
		case <-exitChannel: //监测退出通道

			close(client.C) //关闭用户自带消息通道：以终止messageToClient()

			rwMutex.Lock()                  //全局变量写锁
			delete(clientList, client.addr) //从在线列表将该用户删除
			rwMutex.Unlock()                //解锁

			messageChannle <- makeMsg(client, "logout\n") //广播用户退出消息

			return //结束当前协程
		case <-hasData:
			//如果有数据流通，不做任何操作，计时器归零
		case <-time.After(time.Second * 60):
			//没有数据流通时，全部阻塞60秒后该通道畅通，执行超时退出命令
			close(client.C)
			rwMutex.Lock()
			delete(clientList, client.addr)
			rwMutex.Unlock()
			messageChannle <- makeMsg(client, "time out to leave\n")

			return
		}
	}

}

func main() {

	//创建监听socket
	listerner, err := net.Listen("tcp", "127.0.0.1:8005")
	if err != nil {
		fmt.Println("net.Listen", err)
		return
	}
	defer listerner.Close()

	//创建消息广播goroutine(接收全局消息通道的数据并发送到所有客户端消息通道)
	go messageToAll()

	//循环监听，等待客户端建立连接
	for {
		conn, err := listerner.Accept()
		if err != nil {
			fmt.Println("listerner.Accept", err)
			continue //连接出错时跳过本次继续
		}

		//为每一个客户端创建通信控制goroutine
		go clientControl(conn)
	}

}
