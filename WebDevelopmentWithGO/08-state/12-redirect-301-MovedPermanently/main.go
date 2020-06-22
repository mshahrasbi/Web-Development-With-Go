package main

import (
	"fmt"
	"net/http"
)

func home(res http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at home: ", req.Method, "\n\n")
}

func bar(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar: ", req.Method)
	http.Redirect(res, req, "/", http.StatusMovedPermanently)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/bar", bar)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
