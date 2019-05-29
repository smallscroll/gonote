package main

import "fmt"

func main() {
	//计算2019年3月4号是当前这一年的第几天

	var year, month, day int

	fmt.Println("输入年月日：")
	fmt.Scan(&year, &month, &day)

	days := 0

	for i := 1; i < month; i++ {
		switch i {
		case 1, 3, 5, 7, 8, 10, 12:
			days = 31
		case 4, 6, 9, 11:
			days = 30
		case 2:
			if year%400 == 0 || year%4 == 0 && year%100 != 0 {
				days = 29
			} else {
				days = 28
			}
		}
		day += days
	}

	fmt.Printf("今天是今年的第%d天\n", day)

}
