package main

import (
	"io"
	"net/http"
)

func dog(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src="/resources/toby.jpg">`)
}

func main() {
	http.HandleFunc("/", dog)
	http.Handle("/resources/",
		http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}
