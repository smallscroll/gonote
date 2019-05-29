package main

import "fmt"

// //定义结构体
// type Student struct {
// 	id    int
// 	name  string
// 	score []int //将切片作为结构体成员
// }

// func main() {
// 	/*
// 	   计算学生平均成绩:
// 	   3名学生，不同考试科目，求每名学生的平均成绩
// 	*/

// 	//定义结构体切片
// 	slice := []Student{
// 		//添加切片元素：结构体变量
// 		Student{1001, "张三", []int{10, 9, 9, 6}},
// 		Student{1002, "李四", []int{8, 7}},
// 		Student{1003, "王五", []int{8, 9, 7}},
// 	}
// 	for i := 0; i < len(slice); i++ { //i表示学生
// 		//定义总成绩
// 		sum := 0
// 		for j := 0; j < len(slice[i].score); j++ { //j表示成绩
// 			//计算总成绩
// 			sum += slice[i].score[j]
// 		}
// 		fmt.Printf("姓名：%s，总成绩：%d，平均成绩：%.1f\n", slice[i].name, sum, float64(sum)/float64(len(slice[i].score)))
// 	}

// }

//定义结构体
type Student struct {
	id    int
	name  string
	score map[string]int //各科成绩
}

func main() {
	/*
	   计算学生平均成绩:
	   3名学生，不同考试科目，求每名学生的平均成绩
	*/

	//定义结构体切片
	slice := []Student{
		//添加切片元素：结构体变量
		Student{1001, "张三", map[string]int{"语文": 10, "数学": 9, "英语": 9, "物理": 6}},
		Student{1002, "李四", map[string]int{"数学": 8, "物理": 7}},
		Student{1003, "王五", map[string]int{"语文": 8, "英语": 9, "历史": 7}},
	}
	//循环每一位学生的成绩
	for i := 0; i < len(slice); i++ {
		//定义总成绩
		sum := 0
		for _, v := range slice[i].score {
			sum += v
		}
		fmt.Println(slice[i].name, "平均成绩:", float64(sum)/float64(len(slice[i].score)), slice[i].score)
	}

}
