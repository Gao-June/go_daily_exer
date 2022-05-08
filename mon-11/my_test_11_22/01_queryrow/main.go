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

func queryRowDemo() {
	sqlStr := "select id, name, age from user where id=?"
	// 初始化 user， 之后查询的结果放到 u 里
	var u user

	// 非常重要，确保 QueryRow() 之后调用 Scan()方法，
	// 否则持有的数据库链接不会被释放
	// Scan() 里传指针变量，方便修改
	err := db.QueryRow(sqlStr, 2).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed, err: %v\n", err)
		return
	}

	fmt.Printf("id: %d, name : %s,  age : %d\n", u.id, u.name, u.age)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf(" init db failed!, err : %v\n", err)
	}

	queryRowDemo()
}
