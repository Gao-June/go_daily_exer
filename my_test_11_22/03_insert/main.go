package main

import (
	"database/sql"
	"fmt"

	// 必要！  open() 连接到 SQL
	_ "github.com/go-sql-driver/mysql"
)

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

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf(" init db failed!, err : %v\n", err)
	}

	sqlStr := "insert into user(id, name, age) values (?,?,?)"
	ret, err := db.Exec(sqlStr, 0005, "ma", 221)
	if err != nil {
		fmt.Printf("insert failed, err : %v\n", err)
		return
	}

	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastInsertId failed, err: %v", err)
		return
	}
	fmt.Printf("insert successed!, the id is: %d \n", theID)
}
