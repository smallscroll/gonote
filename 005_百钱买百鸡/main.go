package main

import "fmt"

func main() {
	//百钱买百鸡
	//公鸡：5钱/只，母鸡：3钱/只，小鸡：1钱/3只，用100钱买100只鸡有多少种组合？
	/*
		公鸡最多买20只
		母鸡最多买33只
		小鸡最多买100只
	*/

	for gongji := 0; gongji <= 20; gongji++ {
		for muji := 0; muji <= 33; muji++ {
			xiaoji := 100 - gongji - muji
			if gongji*5+muji*3+xiaoji/3 == 100 && xiaoji%3 == 0 {
				fmt.Printf("公鸡：%d，母鸡：%d，小鸡：%d\n", gongji, muji, xiaoji)
			}
		}
	}

}
