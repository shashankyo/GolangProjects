package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go mysql tutorial")
	db, err := sql.Open("mysql", "root:password@tpc(127.0.0.1:3306)/testdb")

	if err != nil {
		panic(err.Error())

	}

	defer db.Close()

	fmt.Println("successfully connected to my sql database")
}
