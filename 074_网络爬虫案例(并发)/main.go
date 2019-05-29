package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
	豆瓣电影排行榜爬虫

	URL规则：
		https://movie.douban.com/top250?start=0&filter=
		https://movie.douban.com/top250?start=25&filter=
		https://movie.douban.com/top250?start=50&filter=
*/

//封装文件保存函数
func saveTofile(pageIndex int, filmName, filmStar, filmYear [][]string) {
	//按爬取页码创建文件
	f, err := os.Create("/Users/waaa/Downloads/第" + strconv.Itoa(pageIndex) + "页.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	//按字符串写入文件

	//写入标题
	f.WriteString("电影名称,评分,上映年份\n")
	//循环写入全部信息
	for i := 0; i < len(filmName); i++ {
		f.WriteString(filmName[i][1] + "," + filmStar[i][1] + "," + strings.Trim(strings.Trim(filmYear[i][1], "\n"), " ") + "\n")
	}

	fmt.Printf("第%d页保存成功\n", pageIndex)

}

//封装字符串筛选函数
func findStr(str string, pageIndex int, exitChannel chan<- bool) {
	//编译电影名称正则表达式
	reg := regexp.MustCompile(`<img width="100" alt="(?s:(.*?))"`)
	//从字符串中筛选所有电影名称
	filmName := reg.FindAllStringSubmatch(str, -1)

	//编译电影评分正则表达式
	reg = regexp.MustCompile(`<span class="rating_num" property="v:average">(?s:(.*?))</span>`)
	//从字符串中筛选所有电影评分
	filmStar := reg.FindAllStringSubmatch(str, -1)

	//编译电影上映时间正则表达式
	reg = regexp.MustCompile(`<br>(?s:(.*?))&nbsp;/&nbsp;`)
	//从字符串中筛选所有电影上映时间
	filmYear := reg.FindAllStringSubmatch(str, -1)

	//调用文件保存函数
	saveTofile(pageIndex, filmName, filmStar, filmYear)

	//保存完后给退出通道发送信息
	exitChannel <- true

}

//封装网页请求函数
func httpGet(pageIndex int, exitChannel chan<- bool) {

	//初始化URL
	url := "https://movie.douban.com/top250?start=" + strconv.Itoa((pageIndex-1)*25) + "&filter="

	//请求页面，获取响应包
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close() //关闭响应包体

	//读取响应包体的内容并保存到字符串
	buf := make([]byte, 4096)
	var str string
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 { //判断读取结束
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		}
		str += string(buf[:n]) //保存字符串
	}
	//调用字符串筛选函数
	findStr(str, pageIndex, exitChannel)

}

//封装爬取处理函数
func mySpider(startPage int, endPage int) {
	//创建退出通道
	exitChannel := make(chan bool)

	//循环处理所有页面
	for i := startPage; i <= endPage; i++ {
		//并发调用网页请求函数
		go httpGet(i, exitChannel)
	}

	//阻塞接收退出通道的信息，处理完所有页面后该循环才会结束，退出主协程
	for i := startPage; i <= endPage; i++ {
		<-exitChannel
	}

}

func main() {

	//定义爬取起始页和终止页
	var startPage, endPage int

	fmt.Println("输入页面范围：[起始页 终止页]")
	fmt.Scan(&startPage, &endPage)

	//调用爬取处理函数
	mySpider(startPage, endPage)

}
