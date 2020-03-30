package main

import (
	"io"
	"net/http"
)

func dog(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(res, `
		<!--not serving from our server-->
		<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
	`)
}

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}
