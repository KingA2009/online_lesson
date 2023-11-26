package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type User struct {
	Name  string
	Email string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1984"
	dbname   = "postgres"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
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

	var Users []User
	results, err := db.Query("SELECT name, email from users")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var user User
		err = results.Scan(&user.Name, &user.Email)
		if err != nil {
			panic(err.Error())
		}
		Users = append(Users, user)
	}
	fmt.Fprint(w, Users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
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

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]
	_, err = db.Query("insert into users values($1, $2)", name, email)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprint(w, "created succesufly")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
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

	vars := mux.Vars(r)
	name := vars["name"]
	_, err = db.Query("delete from users where name = $1", name)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprint(w, "deleted succesufly")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
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

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]
	_, err = db.Query("update users set email = $1 where name = $2", email, name)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprint(w, "updated succesufly")
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Go ORM Tutorial")

	handleRequest()
}
