// ListenAndServe on port ":8080" using the default ServeMux.

// Use HandleFunc to add the following routes to the default ServeMux:

// "/"
// "/dog/"
// "/me/

// Add a func for each of the routes.

// Have the "/me/" route print out your name.

package main

import (
	"io"
	"net/http"
)

func home(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Home Page")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Dog page")
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "You are in Me Page")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}
