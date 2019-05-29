package main

import "fmt"

func main() {
	//三只小猪称体重
	var a, b, c int

	fmt.Println("分别输入三只小猪的体重：")
	fmt.Scan(&a, &b, &c)

	if a > b {
		if a > c {
			fmt.Println("A重")
		} else {
			fmt.Println("C重")
		}
	} else {
		if b > c {
			fmt.Println("B重")
		} else {
			fmt.Println("C重")
		}
	}
}
