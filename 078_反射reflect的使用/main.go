package main

import (
	"fmt"
	"reflect"
)

/*
	获取变量Kind(类别):

		通过反射，可以拿到变量的kind（类别）。 Reflect.Value.Kind()
		注：Type 是类型 Kind 是类别

	修改变量值:

		借助reflect.Value提供的Set...方法来进行设置。必须通过对应的指针类型来操作，才能改变传入的变量的值。

			reVal.Elem() 相当于：
				var a int = 100
				var p *int = &a
				*p = 222
			reVal.Elem().SetInt(222) 相当于：
				a = *p


*/

//测试1:

func reflectTest(i interface{}) {

	//获取类别
	reType := reflect.TypeOf(i)
	fmt.Printf("%v;", reType.Kind()) //方法1：获得对应变量的具体kind类别，可以和reflect包中的常量列表进行比对

	reVal := reflect.ValueOf(i)
	fmt.Printf("%v\n", reVal.Kind()) //方法2：获得对应变量的具体kind类别，可以和reflect包中的常量列表进行比对

	//判断类别：获取变量值
	if reType.Kind() == reflect.Int {
		n := reVal.Int() + 100 //获取变量值
		fmt.Println("测试1:", n)
	}

	if reType.Kind() == reflect.Bool {
		fmt.Println("测试1:", reVal.Bool)
	}

	if reType.Kind() == reflect.Float64 {
		fmt.Println("测试1:", reVal.Float())
	}

}

//测试2:

func reflectChangeVal(i interface{}) {
	reVal := reflect.ValueOf(i)
	if reVal.Kind() == reflect.Ptr && reVal.Elem().Kind() == reflect.Int {
		reVal.Elem().SetInt(222) //修改变量值
	}

}

//测试3: 获取结构体的值

const tagName = "Testing"

//带有tag标签的结构体类型
type Student struct {
	Name string `Testing:"-"`
	Age  int    `Testing:"stu_age"`
	Id   int    `Testing:"stu_id"`
}

func reflectStruct3(i interface{}) {

	//获取reflect.Type类型，得到interface
	reType := reflect.TypeOf(i)

	//获取reflect.Value类型
	reVal := reflect.ValueOf(i)

	//获取结构体字段数: reType.NumField()
	for i := 0; i < reType.NumField(); i++ { //循环遍历
		structField := reType.Field(i)      //取出每一个字段，得到结构体
		tag := structField.Tag.Get(tagName) //传入结构体成员的tag键，获取tag值

		//打印：字段名、字段类型、字段值、tag值
		fmt.Printf("测试3: %v, %v, %v, %v\n", structField.Name, structField.Type, reVal.Field(i), tag)
	}
}

func main() {

	//测试1: 获取变量值
	var n int = 100
	var bl bool = false
	var flt float64 = 3.14

	reflectTest(n)
	reflectTest(bl)
	reflectTest(flt)

	//测试2: 修改变量值
	reflectChangeVal(&n) //传入地址
	fmt.Println("\n测试2:", n, "\n")

	//测试3: 获取结构体的值
	stu := Student{"jack", 18, 95} //初始化结构体变量
	reflectStruct3(stu)

}
