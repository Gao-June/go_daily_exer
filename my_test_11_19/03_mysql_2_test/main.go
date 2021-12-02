package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "user:passward@tcp(127.0.0.1:3360)/mymytest"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("open err, ", err)
		panic(err)
	}

	fmt.Println("open succedd!")
	defer db.Close()
}
