// 1. Take the previous program in the previous folder and change it so that:
// * a template is parsed and served
// * you pass data into the template

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
	http.HandleFunc("/", home)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
