package main

import (
	"fmt"
	"reflect"
)

/*
	反射操作结构体：

		反射调用结构体方法

*/

//定义结构体类型
type Student struct {
	Id    int    `json:"index"`
	Name  string `json:"name"`
	Age   int    `json:"学员年龄"`
	Score float32
}

//注：方法名首字母必须大写。否则 reflect.Value 的 Method 找不到该方法。

//Student方法1:
func (stu Student) Print() {
	fmt.Println("方法1:", stu)
}

//Student方法2:
func (stu Student) Add(a, b int) int {
	return a + b
}

//Student方法3:
func (stu Student) ResetInfo(id int, name string, age int, score float32) {
	stu.Id = id
	stu.Name = name
	stu.Age = age
	stu.Score = score
}

//使用反射的函数
func useReflect(ref interface{}) {
	reType := reflect.TypeOf(ref) //获取Type
	reVal := reflect.ValueOf(ref) //获取Value

	//类型验证
	reKind := reVal.Kind()
	if reKind != reflect.Struct {
		fmt.Println("参数类型须为结构体")
		return
	}
	//获取字段个数
	numField := reVal.NumField()
	fmt.Println("Student结构体包含字段数：", numField)

	for i := 0; i < numField; i++ {
		tagVal := reType.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("字段%d Val = %v, Tag = %s \n", i+1, reVal.Field(i), tagVal)
		} else {
			fmt.Printf("字段%d Val = %v \n", i+1, reVal.Field(i))
		}
	}

	//获取结构体方法个数
	numMethod := reVal.NumMethod()
	fmt.Printf("Student结构体有 %d 个方法 \n", numMethod)

	//调用一个方法
	//Method方法返回reVal持有值类型的第i个方法的已绑定状态的函数形式的Value封装。使用该函数获取方法时，i的计数从0开始。
	//Call方法使用输入的参数调用reVal持有的函数
	reVal.Method(1).Call(nil)

	//Call方法的参数必须是reflect.Value类型的切片
	var params []reflect.Value
	params = append(params, reflect.ValueOf(39))
	params = append(params, reflect.ValueOf(27))

	sum := reVal.Method(0).Call(params)
	fmt.Println("sum =", sum[0].Int())

}

func main() {
	//初始化结构体变量
	stu := Student{1001, "Jack", 19, 95.5}

	//使用反射的函数
	useReflect(stu)
}
