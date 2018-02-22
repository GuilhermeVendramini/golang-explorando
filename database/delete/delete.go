package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/goexp")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("delete from users where id = ?")
	stmt.Exec(4)
}
