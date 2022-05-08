package main

import (
	"database/sql"
	"fmt"

	// 必要！  open() 连接到 SQL
	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id   int
	name string
	age  int
}

var db *sql.DB

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_go_test"

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// 尝试与数据库建立连接（检验 dsn 是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}

	// 连接成功
	return nil
}

func queryDemo() {
	sqlStr := "select id, name, age from user where id > ?"

	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err: %v \n", err)
		return
	}

	// 非常重要，关闭 rows 释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果 集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err: %v \n", err)
			return
		}

		fmt.Printf("id: %d, name: %s, age : %d\n", u.id, u.name, u.age)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf(" init db failed!, err : %v\n", err)
	}

	queryDemo()
}
