package main

import "fmt"

func main() {

	/*
		统计字符个数
	*/

	str := "yiyiya"
	m := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		m[str[i]]++
	}
	for k, v := range m {
		fmt.Printf("%c出现%d次\n", k, v)
	}
}
