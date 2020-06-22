package main

import (
	"fmt"
	"net/http"
)

func foo(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)

	if req.URL.Path != "/" {
		http.NotFound(res, req)
	}

	// fmt.Println(req.URL.Path)
	fmt.Fprintln(res, "go look at your terminal")
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}
