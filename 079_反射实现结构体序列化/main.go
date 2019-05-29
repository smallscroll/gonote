package main

import (
	"fmt"
	"reflect"
	"strings"
)

/*
	反射操作结构体：

		反射实现结构体序列化：

		Person包含四个字段，使用一个名为lable的tag来定义打印时的标题。
		如果是字符串类型的字段，通过名为uppercase的tag控制是否显示字符串的大写形式（默认小写）。


*/

//结构体类型：
//要求：除 TOM 名字外，其余字段，无论输入为何，输出皆为小写。
type Person struct {
	Name        string `label:"Person Name: " uppercase:"true"` //包含两个tag
	Age         int    `label:"Age is: "`
	Sex         string `label:"Sex is: "`
	Description string
}

func myMarshl(iptr interface{}) {

	//获取传入参数的Type类型
	reType := reflect.TypeOf(iptr)

	//判断传入参数类型是否正确
	if reType.Kind() != reflect.Ptr || reType.Elem().Kind() != reflect.Struct {
		fmt.Println("请传入一个结构体指针")
		return
	}

	//获取传入参数的Value类型并取指针指向的结构体变量
	reVal := reflect.ValueOf(iptr).Elem()

	//获取结构体字段个数
	for i := 0; i < reVal.NumField(); i++ {
		structField := reVal.Type().Field(i) //获取结构体字段信息
		tag := structField.Tag               //获取字段tag(key:value)
		label := tag.Get("label")            //根据tag的键获取tag的值
		//判断标签是否为空，空则使用字段名
		if label == "" {
			label = structField.Name + ": "
		}
		//获取所有字段的值，将其保存为string
		value := fmt.Sprintf("%v", reVal.Field(i))

		//借助Type的Kind判断字段类型是否为string
		if structField.Type.Kind() == reflect.String {
			//字段类型为字符串时接着判断其是否规定大小写
			if tag.Get("uppercase") == "true" {
				value = strings.ToUpper(value) //
			} else {
				value = strings.ToLower(value)
			}
		}
		//拼接输出结果
		fmt.Println(label + value)
	}
}

func main() {

	//初始化结构体变量
	person := Person{"Tom", 29, "Male", "Cool"}
	myMarshl(&person) //传入地址

}
