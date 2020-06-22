package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func index(res http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(res, "Successfully completed.")
	check(err)
}

func main() {
	// root:<password>@localhost:3306/node_complete?charset=utf8
	db, err = sql.Open("mysql", "root:<!!!!!>@tcp(localhost:3306)/node_complete?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)
}
