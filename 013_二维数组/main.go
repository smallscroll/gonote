package main

import "fmt"

/*
	二维数组
*/

func main() {

	//定一个2行3列的二维数组
	var arr [2][3]int = [2][3]int{{11, 12, 13}, {21, 22, 23}}

	fmt.Println(len(arr))    //打印二维数组的长度：结果为它的行数
	fmt.Println(len(arr[0])) //打印二维数组索引为0行的长度：结果为它的列数（元素个数）

	//循环打印二维数组的值
	//外层循环控制行
	for i := 0; i < len(arr); i++ {
		//内层循环控制列
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println("")
	}

	fmt.Println(arr[1][1]) //打印二维数组行索引为1列索引为1的元素值

	/*
		三维数组
	*/

	fmt.Println("三维数组：")

	//定义一个2层2行3列的三维数组
	var arr2 [2][2][3]int = [2][2][3]int{
		{
			{111, 112, 113},
			{121, 122, 123},
		},
		{
			{211, 212, 213},
			{221, 222, 223},
		},
	}
	//循环打印三维数组的值
	for l := 0; l < len(arr2); l++ {
		for m := 0; m < len(arr2[l]); m++ {
			for n := 0; n < len(arr2[l][m]); n++ {
				fmt.Print(arr2[l][m][n], " ")
			}
			fmt.Println("")
		}
		fmt.Println("")
	}

}
