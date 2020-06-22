package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hi doggy!")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hi kitty!")
}

func main() {
	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)

	// or
	http.Handle("/dog", http.HandleFunc(d))
	http.Handle("/cat", http.HandleFunc(c))

	http.ListenAndServe(":8080", nil)
}
