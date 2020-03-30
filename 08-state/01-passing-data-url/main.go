package main

import (
	"fmt"
	"net/http"
)

func foo(res http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	fmt.Fprintln(res, "Do my serach: "+v)
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

// usage: localhost:8080/?q=dog
