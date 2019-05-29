package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	/*
		字符串查找
	*/

	str1 := "hello, world."
	str2 := "hello"

	//模糊查找：判断str2是否包含在str1中，返回一个布尔值
	if strings.Contains(str1, str2) {
		fmt.Println("找到了")
	} else {
		fmt.Println("没找到")
	}

	//位置查找：查找字符串2在字符串1中的位置，找到返回一个下标，未找到返回值为-1
	index := strings.Index(str1, str2)
	fmt.Println(index)

	/*
		字符串切片拼接
	*/

	slice1 := []string{"hello", "world", "你好", "哈哈哈"}
	//用“—”符号作为连接符将字符串切片进行拼接，返回一个字符串
	str3 := strings.Join(slice1, "-")
	fmt.Println(str3)

	/*
		字符串重复
	*/
	//返回一个重复2次的字符串
	ch1 := strings.Repeat(str3, 2)
	fmt.Println(ch1)

	/*
		字符串内容替换
	*/

	//将字符串3中的“o”替换为“0”，替换2次，返回一个字符串：
	//（字符串，旧字符，新字符，替换次数）替换次数为-1表示全部替换
	ch2 := strings.Replace(str3, "o", "0", 2)
	fmt.Println(ch2)

	/*
		字符串切割
	*/

	//把字符串3中的“-”作为分割符进行切割，返回值为一个字符串切片
	ch3 := strings.Split(str3, "-")
	fmt.Println(ch3)

	/*
		字符串首尾字符去除
	*/

	str4 := "     哈哈 哈哈哈 "
	//把字符串4首尾的空格去除，返回一个字符串
	ch4 := strings.Trim(str4, " ")
	fmt.Println(ch4)

	/*
		字符串空格去除并切割
	*/

	str5 := "ni hao ya ha ha ha"
	//去掉字符串5中的空格并进行分割，返回值为一个字符串切片
	ch5 := strings.Fields(str5)
	fmt.Printf("%T\n\n", ch5)

	/*
		判断字符串前缀和后缀
	*/
	//判断后缀：返回值为bool类型
	bl := strings.HasSuffix("he.jpg", ".jpg")
	fmt.Println(bl)
	//判断前缀
	bl = strings.HasPrefix("he.jpg", "he")
	fmt.Println(bl)

	/*
		字符串的类型转换
	*/

	//
	//其他类型转换为字符串(Format)
	//

	//布尔类型转换为字符串
	b := true
	str6 := strconv.FormatBool(b)
	fmt.Printf("%q\n", str6)

	//整型数字按10进制转换为字符串
	str7 := strconv.FormatInt(12335, 10)
	fmt.Printf("%q\n", str7)
	//十进制数字快速转换为字符串
	str8 := strconv.Itoa(123456)
	fmt.Printf("%q\n", str8)

	//浮点型转换成字符串：（数字，类型，保留小数位，处理位数）
	str9 := strconv.FormatFloat(3.1415, 'f', 3, 64)
	fmt.Printf("%q\n", str9)

	//
	//字符串转换为其他类型(Parse)
	//

	//将字符串转换为布尔值
	b2, _ := strconv.ParseBool("ture")
	fmt.Println(b2)

	//字符串转换为整型（字符串，进制，处理位数）
	v1, _ := strconv.ParseInt("1234567", 10, 64)
	fmt.Println(v1)
	//字符串快速转换为十进制整型
	v2, _ := strconv.Atoi("12345678")
	fmt.Println(v2)

	//字符串转换为浮点型（字符串，处理位数）
	f1, _ := strconv.ParseFloat("3.14159", 64)
	fmt.Println(f1)

	fmt.Println("")

	/*
		其他类型转换为字符切片(Append)
	*/

	//定义一个空字符切片
	var slice2 []byte
	//布尔类型转换为字符切片
	slice2 = strconv.AppendBool(slice2, true)
	fmt.Println(string(slice2))
	//整型转换为字符切片（切片，整型，进制）
	slice2 = strconv.AppendInt(slice2, 123456789, 10)
	fmt.Printf("%c\n", slice2)
	//浮点型转换为字符切片（切片，浮点数，类型，保留小数位，处理位数）
	slice2 = strconv.AppendFloat(slice2, 3.14159, 'f', 3, 64)
	fmt.Printf("%c\n", slice2)
	//字符串转换为字符切片(带引号一起转换)
	slice2 = strconv.AppendQuote(slice2, "hello")
	fmt.Printf("%c\n", slice2)
	fmt.Println(len(slice2))
	fmt.Println(string(slice2)) //强制将字符切片转换为字符串
}
