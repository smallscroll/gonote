package main

import (
	"fmt"
	"reflect"
)

/*
	reflect
	反射是一种运行时的机制，能够在运行时更新、检查变量值，调用方法。
	reflect包实现了运行时反射，允许程序操作任意类型的对象。(这里面有两个核心的类型Type和Value)


		• 反射可以在运行时动态的获取变量信息。如变量类型（type）、类别（kind）
		• 反射可以在运行时获取结构体本身的信息，如：字段、方法等信息。
		• 反射可以在运行时动态修改变量值，调用关联方法。


	反射使用的基本原理：
		变量 -> TypeOf() -> Type类型(interface) -> interface{ Kind(),Name(),Method()... }
		变量 -> ValueOf() -> Value类型(struct) -> 可以调用{ Kind(),Type(),Name(),call()... }

		借助这些方法完成变量到Type，通过Type反过来再操作变量。或者从变量到Value类型，通过Value提供的方法反向再操作这个变量。


	反射的应用：

		基础类型转换：
			方法1: 变量 -> inerface -> ValueOf() -> v:=reflect.Value -> v.Interface() -> interface -> 类型断言 -> 变量
			方法2: 变量 -> inerface -> ValueOf() -> v:=reflect.Value ->Int(),Bool()...运算
					(只针对基础数据类型，对自定义类型如结构体不适用)


*/

//测试1: 封装一个interface{}作为参数的函数，专门用于做反射

func reflectTest(i interface{}) {

	//得到reflect.Value类型
	reVal := reflect.ValueOf(i) //使用 reflect.ValueOf() 函数将 interface{} 转成 reflect.Value 类型
	fmt.Printf("%T\n", reVal)   //reflect.Value

	//reflect.Value类型转换为interface
	iVal := reVal.Interface()         //返回reVal当前持有的值
	fmt.Printf("%T,%v\n", iVal, iVal) //int,100 //不能进行算术运算

	//iVal通过类型断言转换为int类型
	val := iVal.(int)                    //使用类型断言，切实的指向一个具体的数据类型
	fmt.Printf("测试1: %T,%v\n", val, val) //int,100 //能进行算术运算

}

//测试2:

func reflectTest2(i interface{}) {
	reVal := reflect.ValueOf(i)
	n := reVal.Int() + 200 //调用Int()方法直接进行转换（针对基础数据类型）
	fmt.Println("测试2:", n)
}

//测试3:

type Student struct {
	Name string
	Age  int
	Id   int
}

//封装函数完成类型转换
func reflectStruct3(i interface{}) {
	//获取Type类型
	reType := reflect.TypeOf(i)
	fmt.Printf("%v\n", reType)

	//获取Value类型
	reVal := reflect.ValueOf(i)
	fmt.Printf("%v,%T\n", reVal, reVal) //{jack 18 95},reflect.Value //不能直接使用

	//将reVal转换为interface
	ival := reVal.Interface()
	fmt.Printf("%v,%T\n", ival, ival)

	//ival接口断言,添加判断机制
	if val, ok := ival.(Student); ok {
		fmt.Println("测试3:", val.Name)
	}
}

func main() {

	//定义一个int类型变量
	var num int = 100

	//测试1: 使用反射函数（变量 -> interface）
	reflectTest(num)

	//测试2: 使用反射直接获取变量值进行运算
	reflectTest2(num)

	//测试3: 结构体类型转换
	stu := Student{"jack", 18, 95}
	reflectStruct3(stu)

}
