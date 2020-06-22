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

func home(res http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at home: ", req.Method, "\n\n")
}

func bar(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar: ", req.Method)

	// process form submission here

	http.Redirect(res, req, "/", http.StatusSeeOther)
}

func barred(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred: ", req.Method)
	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
