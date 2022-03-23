package main

import (
	"fmt"
	"gonote/085_Protobuf基本操作/pb"

	"github.com/golang/protobuf/proto"
)

/*
	使用protobuf协议之前须创建proto文件并将其编译为go文件：
		示例：将pb目录下创建的所有.proto文件进行编译并将生成的go文件输出到pb目录
			$cd pb
			$protoc --go_out=. *.proto
*/

func main() {

	/*
		编码：
	*/

	//初始化一个protobuf结构体对象（继承自.proto文件里的message）
	stu := &pb.Student{
		Name:   "Jack",
		Age:    18,
		Emails: []string{"crux@waaa.org", "club@waaa.org"},
		Phones: []*pb.PhoneNumber{
			&pb.PhoneNumber{
				Number: "18600000000",
			},
			&pb.PhoneNumber{
				Number: "13900000000",
			},
		},
		Data: &pb.Student_Socre{ //给实现了sStudent_Data接口的Student_Socre结构体对象进行赋值
			Socre: 100,
		},
	}

	//将protobuf结构体对象转化为二进制数据
	data, err := proto.Marshal(stu)
	if err != nil {
		fmt.Println("proto.Marshal error", err)
		return
	}

	/*
		解码：
	*/

	//定义一个protobuf结构体对象
	newStu := &pb.Student{}

	//将protobuf数据解码到protobuf结构体对象中
	err = proto.Unmarshal(data, newStu)
	if err != nil {
		fmt.Println("proto.Unmarshal error", err)
		return
	}
	fmt.Println(newStu)
	fmt.Println("使用实现了接口的结构体对象方法取值：", newStu.GetSocre())
}
