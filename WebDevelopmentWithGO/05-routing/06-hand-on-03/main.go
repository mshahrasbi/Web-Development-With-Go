// 1. Take the previous program and change it so that:
// * func main uses http.Handle instead of http.HandleFunc

// Contstraint: Do not change anything outside of func main

// Hints:

// [http.HandlerFunc](https://godoc.org/net/http#HandlerFunc)
// ``` Go
// type HandlerFunc func(ResponseWriter, *Request)
// ```

// [http.HandleFunc](https://godoc.org/net/http#HandleFunc)
// ``` Go
// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
// ```

// [source code for HandleFunc](https://golang.org/src/net/http/server.go#L2069)
// ``` Go
//   func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
//   		mux.Handle(pattern, HandlerFunc(handler))
//   }
// ```

package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Home page")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Dog Page")
}

func me(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("something.gohtml")

	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(w, "something.gohtml", "John Smith")

	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func main() {
	http.Handle("/", http.HandleFunc(home))
	http.Handle("/dog/", http.HandleFunc(dog))
	http.Handle("/me/", http.HandleFunc(me))

	http.ListenAndServe(":8080", nil)
}
