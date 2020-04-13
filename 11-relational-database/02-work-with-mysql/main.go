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
	_, err := io.WriteString(res, "at index")
	check(err)

}

func amigos(res http.ResponseWriter, req *http.Request) {

	rows, err := db.Query(`SELECT aName FROM amigos`)
	check(err)
	defer rows.Close()

	// data to be used in query
	var s, name string
	s = "RETRIEVED RECORDS: \n"

	// query
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}

	fmt.Fprintln(res, s)

}

func create(res http.ResponseWriter, req *http.Request) {

	stmt, err := db.Prepare(`CREATE TABLE customer (name VARCHAR(20));`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(res, "CREATE TABLE customer", n)
}

func insert(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO customer VALUES("jAMES");`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(res, "INSERTED RECORD", n)
}

func read(res http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT * FROM customer;`)
	check(err)
	defer rows.Close()

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		fmt.Fprintln(res, "RETRIEVED RECORD: ", name)
	}
}

func update(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`UPDATE customer SET name="Jimmy" WHERE name="James";`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(res, "UPDATED RECORD", n)

}

func del(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM customer WHERE name="Jimmy";`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(res, "DELETED RECORD", n)
}

func drop(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DROP TABLE customer;`)
	check(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(res, "DROPPED TABLE customer")
}

func main() {
	// root:<password>@localhost:3306/node_complete?charset=utf8
	db, err = sql.Open("mysql", "root:<!!!!!!!>@tcp(localhost:3306)/node_complete?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", del)
	http.HandleFunc("/drop", drop)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	err := http.ListenAndServe(":8080", nil)
	check(err)
}
