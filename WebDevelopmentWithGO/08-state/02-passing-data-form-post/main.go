package main

import (
	"io"
	"net/http"
)

func foo(res http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	// post it is go through body
	io.WriteString(res, `
		<form method="POST">
			<input type="text" name="q">
			<input type="submit">
		</form>
		<br>
	`+v)
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
