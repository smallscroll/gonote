package main

import "fmt"

//定义接口
type Calcer interface {
	Result() int
}

//操作父类
type Operate struct {
	n1 int
	n2 int
}

//加法子类
type AddOpt struct {
	Operate
}

//减法子类
type SubOpt struct {
	Operate
}

//乘法子类
type MulOpt struct {
	Operate
}

//除法子类
type DivOpt struct {
	Operate
}

//加法方法
func (a *AddOpt) Result() int {
	return a.n1 + a.n2
}

//减法方法
func (s *SubOpt) Result() int {
	return s.n1 - s.n2
}

//乘法方法
func (m *MulOpt) Result() int {
	return m.n1 * m.n2
}

//除法方法
func (d *DivOpt) Result() int {
	return d.n1 / d.n2
}

//多态实现*

func Result(calc Calcer) {
	value := calc.Result() //通过接口调用对象方法
	fmt.Println(value)
}

/*
	工厂设计模式
*/

//创建工厂类
type Factory struct {
}

//创建工厂方法
func (f *Factory) CreateFactory(num1 int, num2 int, ch string) {
	switch ch {
	case "+":
		//创建加法对象
		add := AddOpt{Operate{num1, num2}}
		//通过多态调用加法方法
		Result(&add)
	case "-":
		sub := SubOpt{Operate{num1, num2}}
		Result(&sub)
	case "*":
		mul := MulOpt{Operate{num1, num2}}
		Result(&mul)
	case "/":
		div := DivOpt{Operate{num1, num2}}
		Result(&div)
	}

}

func main() {

	// //实例化加法对象
	// add := AddOpt{Operate{10, 20}}

	// //实例化乘法对象
	// mul := MulOpt{Operate{2, 3}}

	// //通过多态调用对象方法*
	// Result(&add)
	// Result(&mul)

	/*
		工厂设计模式
	*/

	//创建工厂对象
	var f Factory
	//传递数据并调用工厂方法
	f.CreateFactory(100, 200, "+")
	f.CreateFactory(10, 3, "-")
	f.CreateFactory(2, 3, "*")
	f.CreateFactory(10, 2, "/")

}
