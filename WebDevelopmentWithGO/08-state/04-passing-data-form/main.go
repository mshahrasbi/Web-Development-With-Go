package main

import (
	"html/template"
	"log"
	"net/http"
)

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func indexPage(res http.ResponseWriter, req *http.Request) {

	f := req.FormValue("first")
	l := req.FormValue("last")
	s := req.FormValue("subscribe") == "on"

	err := tpl.ExecuteTemplate(res, "index.gohtml", person{f, l, s})
	if err != nil {
		http.Error(res, err.Error(), 500)
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", indexPage)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
