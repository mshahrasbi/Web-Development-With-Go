package main

import (
	"io"
	"net/http"
	"os"
)

func dog(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(res, `
		<img src="/toby.jpg">
	`)
}

func dogPic(res http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg")

	if err != nil {
		http.Error(res, "file not found", 404)
		return
	}

	defer f.Close()

	io.Copy(res, f)
}

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)

	http.ListenAndServe(":8080", nil)
}
