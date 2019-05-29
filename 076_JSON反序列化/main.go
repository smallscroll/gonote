package main

import (
	"encoding/json"
	"fmt"
)

/*
	JSON反序列化：
		反序列化json字符串时，务必确保反序列化传出的数据类型，与之前序列化的数据类型完全一致。

*/

//定义一个结构体
type Student struct {
	Name  string `json:"stu_name"`
	Id    int
	Age   int
	Score float64 `json:"sru_score,string"`
}

//结构体反序列化函数
func UnserialStruct() {

	//定义一个结构体变量
	var stu Student

	//JSON字符串
	str := `{"stu_name":"Yoyo","Id":1001,"Age":18,"sru_score":"98.5"}`

	//反序列化
	err := json.Unmarshal([]byte(str), &stu)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(stu) //打印反序列化后数据
}

//Map反序列化函数
func UnserialMap() {

	//定义并初始化一个map空间
	var m map[string]interface{}

	//JSON字符串
	str := `{"age":19,"food":["苹果","菠萝","芒果"],"name":"小明","salary":155.5}`

	//反序列化
	//map反序列化时会自动开辟内存空间，但是传参时必须引用传递
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m) //打印反序列化后数据
}

//切片反序列化函数
func UnserialSlice() {

	//定义一个map切片
	var s []map[string]interface{}

	//JSON字符串
	str := `[{"age":18,"food":["苹果","菠萝","芒果"],"name":"小明","salary":155.5}]`

	//反序列化
	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s) //打印反序列化后数据
}

func main() {

	//结构体反序列化
	UnserialStruct()

	//Map反序列化
	UnserialMap()

	//切片反序列化
	UnserialSlice()

}
