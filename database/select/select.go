package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type users struct {
	id   int
	name string
}

func main() {
	db, err := sql.Open("mysql", "root:root@/goexp")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, name from users where id < ?", 3)
	defer rows.Close()

	for rows.Next() {
		var u users
		rows.Scan(&u.id, &u.name) //Atribuindo o resultado a estrutura users
		fmt.Println(u)
	}
}
