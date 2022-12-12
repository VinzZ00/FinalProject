package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	userId   int
	userName string
	email    string
	passkey  string
}

func main() {
	fmt.Println("Hello wolrd")

	db, err := sql.Open("mysql", "root@/bncc_go")

	if err != nil {
		panic(err)
	}

	fmt.Println(db)

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	stmt, err := db.Prepare("select * from msUser where userId = ?")

	if err != nil {
		panic(err)
	}

	var u User

	err = stmt.QueryRow("2").Scan(&u.userId, &u.userName, &u.email, &u.passkey)

	if err != nil {
		panic(err)
	}

	fmt.Println(u)

}
