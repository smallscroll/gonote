package main

import (
	"encoding/json"
	"fmt"
)

/*
	JSON

		序列化：将数据类型转换为JSON字符串
		反序列化：将JSON字符串转回成对应的数据类型

	JSON语法：
		1.键值对：
				key: 必须是字符串
				value: 可存放对象{}、数组[]、数字、字符串、逻辑值、空
		2.对象{}：可放键值对、数组[]、对象{}
		3.数组[]：可放键值对、数组[]、对象{}
*/

//定义一个结构体（该结构体使用了`结构体标签`：tag）
type Student struct {
	Name  string  `json:"stu_name"`          //通过结构体标签指定JSON序列化时键的名称
	Id    int     `json:"-"`                 //当tag值指定为"-"时，序列化时会忽略处理该tag
	Age   int     `json:"stu_age,omitempty"` //指定名称并在序列化时忽略0值和空值
	Score float64 `json:"sru_score,string"`  //指定名称和序列化后的数据类型
}

func main() {
	stu := Student{"Jack", 1001, 18, 98.5}

	/*
		结构体序列化
	*/

	//JSON序列化(参数：空接口类型意味着任何数据类型都可以序列化) 返回值：字符切片,err
	//结构体成员变量进行序列化时成员首字母需要大写(外部的JSON包调用)
	//当结构体成员使用了`结构体标签`时，序列化使用标签指定的参数进行序列化
	data1, err := json.Marshal(stu)
	if err != nil {
		fmt.Println(err)
		return
	}
	//打印JSON字符串
	fmt.Println("结构体序列化：", string(data1))

	/*
		Map序列化
	*/

	//定义并初始化一个map空间
	m := make(map[string]interface{})

	m["name"] = "小明"
	m["age"] = 18
	m["salary"] = 155.5
	m["food"] = [3]string{"苹果", "菠萝", "芒果"}

	//JSON序列化
	data2, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	//打印JSON字符串
	fmt.Println("Map序列化：", string(data2))

	/*
		切片序列化
	*/

	//定义一个map切片
	var s []map[string]interface{}
	//将m追加到切片
	s = append(s, m)

	//JSON序列化
	data3, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	//打印JSON字符串
	fmt.Println("切片序列化：", string(data3))

}
