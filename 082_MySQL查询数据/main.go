package main

/*
	查询数据
*/

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //导入数据库驱动
)

func main() {

	/*
		测试连接数据库
	*/

	//测试连接数据库，返回一个DB指针和error
	db, err := sql.Open("mysql", "test:123456@tcp(127.0.0.1)/test01")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	/*
		连接数据库
	*/

	//Ping()方法函数调用成功后，DB则指向一个连接成功的数据库对象
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	/*
		单行查询

	*/

	//使用函数QueryRow()读取数据库表中一条记录。即使写入的sql语句能查询多条结果，该函数也只返回一条记录。

	//组织SQL语句
	sql := ` select * from st7 where id=4`

	//根据SQL查询字段数定义接收数据的变量
	var id, name string

	//调用单行查询函数并执行SQL语句，将返回的查询结果保存到row中
	row := db.QueryRow(sql)

	//使用Scan()方法提取字段数据（取地址）
	err = row.Scan(&id, &name)
	if err != nil {
		fmt.Println(err)
		return
	}

	//打印数据
	fmt.Println(id, name)

	/*
		多行查询
	*/
	fmt.Println("\n多行查询:")

	//使用函数Query()读取数据库表中多条记录。

	//组织SQL语句
	sql = `select * from st7;`

	//调用多行查询函数并执行SQL语句，将返回的查询结果保存到rows中
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	//for循环判断游标是否指向一条记录，使用函数Next()调用游标（读取一行后游标自动转到下一行）
	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}

	/*
		预处理查询
	*/
	fmt.Println("\n预处理查询:")

	//组织带有占位符?的SQL语句
	sql = `select * from st7 where id >= ?`

	//传入SQL语句并执行预处理
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Println(err)
		return
	}
	//使用预处理对象调用Query()，同时传入占位符的数据并执行SQL语句
	rows, err = stmt.Query(5)
	if err != nil {
		fmt.Println(err)
		return
	}
	//for循环判断游标是否指向一条记录，使用函数Next()调用游标
	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
}
