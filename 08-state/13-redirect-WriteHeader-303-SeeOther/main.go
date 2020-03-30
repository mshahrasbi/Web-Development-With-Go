package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func home(res http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")
}

func bar(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar:", req.Method)
	// process form data
	res.Header().Set("Location", "/")
	res.WriteHeader(http.StatusSeeOther)
}

func barred(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred:", req.Method)
	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}
