package main

import (
	"database/sql"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	db, err := sql.Open("mysql", "root:root1234@tcp(localhost:3306)/pro-Test")
	if err != nil {
		fmt.Println("error occurred. err=", err)
	}
	defer db.Close()
}
