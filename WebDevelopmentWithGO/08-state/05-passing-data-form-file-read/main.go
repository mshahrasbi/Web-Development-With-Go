package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request) {
	var s string

	fmt.Println(req.Method)

	if req.Method == http.MethodPost {
		// open
		f, h, err := req.FormFile("q")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		defer f.Close()

		// for your information
		fmt.Println("\nfile: ", f, "\nheader: ", h, "\nerror: ", err)

		// read
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		s = string(bs)
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `
		<form method="POST" enctype="multipart/form-data">
			<input type="file" name="q">
			<input type="submit">
		</form>
	`+s)
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
