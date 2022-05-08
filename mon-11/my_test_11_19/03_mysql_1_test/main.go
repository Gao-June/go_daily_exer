package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 数据库信息
	// 默认ip是 127.0.0.1 , 默认端口是 3306
	// dbname 是库名
	dsn := "root:123456@tcp(127.0.0.1:3306)/mymytest"

	// 连接数据库
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("dsn:%s invalid, err: %v \n", dsn, err)
		return
	}
	err = db.Ping() // 尝试连接数据库
	if err != nil {
		fmt.Printf("open %s failed, err: %v \n", dsn, err)
		return
	}
	fmt.Println("连接数据库成功")
}
