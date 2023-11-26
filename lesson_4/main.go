package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1984"
	dbname   = "postgres"
)

type User struct {
	name string `json:"name"`
}

func main() {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	results, err := db.Query("SELECT name from users")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var user User

		err = results.Scan(&user.name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.name)
	}
}
