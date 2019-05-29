package main

/*
	驱动检查：

		查看当前主机是否已安装MySQL驱动：
			sudo find ~/ -name "go-sql-driver"
		如果没有，安装驱动：
			go get github.com/go-sql-driver/mysql
		查看驱动：
			cd ~/go/src/github.com/go-sql-driver

*/

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //导入数据库驱动包，并执行它里面的所有init()函数（程序在操作数据库的时候只需要用到database/sql，而不需要直接使用数据库驱动，所以将这个包的名字设置成下划线。）
)

func main() {

	/*
		测试连接数据库
	*/

	//测试连接数据库，返回一个DB指针和error
	//参数1: driverName 驱动名。"是数据库驱动注册到 database/sql 时所使用的名字"
	//参数2: dataSourceName 数据库连接信息语法。"用户名:密码@连接方式(主机名:端口号)/数据库名?参数"
	db, err := sql.Open("mysql", "test:123456@tcp(127.0.0.1:3306)/test01?loc=Local&parseTime=true") //设置数据库的时区为本地时区并西东解析数据库时间格式
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	//sql.Open不会建立数据库连接, 也不会对数据库链接参数的合法性做检验,
	//它仅仅是初始化一个sql.DB对象，但如果没有导入正确的驱动，该函数会报错。

	/*
		连接数据库
	*/

	//Ping()方法函数调用成功后，DB则指向一个连接成功的数据库对象
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("恭喜！mysql 数据库连接成功!")

	/*
		DB的主要方法有：
			Query 执行数据库的Query操作，例如一个Select语句，返回*Rows
			QueryRow 执行数据库至多返回1行的Query操作，返回*Row
			PrePare 准备一个数据库query操作，返回一个*Stmt，用于后续query或执行。这个Stmt可以被多次执行，或者并发执行
			Exec 执行数不返回任何rows的据库语句，例如delete操作
	*/

	/*
		插入单行数据

		提前在数据库中准备测试表：
		create table st7(
			id int primary key,
			name varchar(10)
		);
	*/

	//使用 DB 的 Exec() 函数来完成SQL语句在go程序中的执行，访问数据库表

	//组织sql语句
	sql := "insert into st7 values(1, 'tom');"

	//执行sql语句
	result, err := db.Exec(sql)
	if err != nil {
		fmt.Println(err)
		return
	}
	n, _ := result.RowsAffected() //获取SQL语句影响的行数

	fmt.Printf("insert into ok! %d row affected.\n", n)

	/*
		插入多行数据
	*/

	//组织SQL语句
	sql = `insert into st7 values(3, 'rose'), (4, '李白'), (5, '杜甫');`
	result, err = db.Exec(sql)
	if err != nil {
		fmt.Println(err)
		return
	}
	n, _ = result.RowsAffected()
	fmt.Printf("insert into ok! %d row affected.\n", n)

	/*
		预处理插入多行数据
	*/

	//准备要插入的数据
	str := [][]string{{"6", "李清照"}, {"7", "辛弃疾"}, {"8", "陆游"}}

	//使用DB 的 Prepare()函数，执行带有预处理的SQL语句

	//组织带占位符?的SQL语句并传入预处理函数
	stmt, err := db.Prepare("insert into st7 values(?,?)")
	if err != nil {
		fmt.Println("Prepare err:", err)
		return
	}
	//遍历str数据，传入参数替换预处理语句中的占位符?并执行
	for _, s := range str {
		stmt.Exec(s[0], s[1]) //传入参数并执行SQL
	}

}
