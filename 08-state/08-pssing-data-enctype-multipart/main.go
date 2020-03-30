package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func index(res http.ResponseWriter, req *http.Request) {
	// body
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	body := string(bs)

	err := tpl.ExecuteTemplate(res, "index.gohtml", body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
