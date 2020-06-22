package main

import (
	"io"
	"net/http"
)

func dog(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(res, `
		<!--Image doesn't serve, not going to display it-->
		<img src="/toby.jpg">
	`)
}

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}
