package main

import "fmt"

func main() {
	//十只小猪称体重
	pigs := [10]int{4, 5, 7, 8, 7, 3, 4, 9, 4, 3}
	//定义最大和最小的初始值
	max, min := pigs[0], pigs[0]
	//定义小猪重量总和
	sum := 0
	//外层循环1次
	for i := 0; i < len(pigs); i++ {
		//内层循环1周，依次把每一只小猪和最大值比较
		if pigs[i] > max {
			//如果体重比最大初始值大则修改最大值
			max = pigs[i]
		}
		//内层循环1周，依次把每一只小猪和最小值比较
		if pigs[i] < min {
			//如果体重比最小初始值小则修改最小值
			min = pigs[i]
		}
		//将所有小猪的值相加
		sum += pigs[i]
	}
	//打印记结果
	fmt.Printf("最重值：%d，最轻值：%d，平均值：%.1f\n", max, min, float64(sum)/float64(len(pigs)))

}
