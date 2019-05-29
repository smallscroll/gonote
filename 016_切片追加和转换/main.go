package main

import "fmt"

func main() {

	/*
		切片追加
	*/

	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{6, 7, 8, 9, 10}
	//使用不定参格式...向一个切片追加另一个切片
	s1 = append(s1, s2...)
	fmt.Println(s1)
	//使用范围遍历追加
	s3 := []int{11, 12, 13}
	for _, v := range s3 {
		s1 = append(s1, v)
	}
	fmt.Println(s1)

	/*
		字符串和字符切片转换
	*/
	fmt.Println("字符串转换为字符切片：")
	str1 := "hello"
	s4 := []byte(str1)
	fmt.Println(s4) // 打印结果为ASICII码
	//遍历字符切片
	for _, v := range s4 {
		fmt.Printf("%c", v)
	}

	fmt.Println(" ")
	fmt.Println("字符切片转换为字符串：")
	s5 := []byte{'h', 'e', 'l', 'l', 'o'}
	str2 := string(s5)
	fmt.Println(str2)
	//字符串中len()表示有效字符个数 字符串结束标志为'\0'
	fmt.Println(len(str2))

	/*
		中文字符串
	*/

	fmt.Println("中文字符串：")
	str6 := "你好"
	for i := 0; i < len(str6); i++ {
		fmt.Println(str6[i]) //6个ASCII码
	}
	//GO语言中一个中文占3个字节
	fmt.Println(len(str6))
	for _, v := range str6 {
		fmt.Printf("%c", v)
	}

}
