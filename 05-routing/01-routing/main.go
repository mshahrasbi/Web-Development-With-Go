package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "Hi doggy!")
	case "/cat":
		io.WriteString(w, "Hi kitty!")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
