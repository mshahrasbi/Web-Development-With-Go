// # Starting with the code in the "starting-files" folder:

// ## wire this program up so that it works

// ParseGlob in an init function

// Use HandleFunc for each of the routes

// Combine apply & applyProcess into one func called "apply"

// Inside the func "apply", use this code to create the logic to respond differently to a POST method and a GET method

// ``` go
// if req.Method == http.MethodPost {
//     		// code here
//     		return
//     	}
// ```

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

func index(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "index.gohtml", nil)
	HandleError(res, err)
}

func about(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "about.gohtml", nil)
	HandleError(res, err)
}

func contact(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "contact.gohtml", nil)
	HandleError(res, err)
}

func apply(res http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {
		err := tpl.ExecuteTemplate(res, "applyProcess.gohtml", nil)
		HandleError(res, err)
		return
	}

	err := tpl.ExecuteTemplate(res, "apply.gohtml", nil)
	HandleError(res, err)
}

func HandleError(res http.ResponseWriter, err error) {
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about/", about)
	http.HandleFunc("/contact/", contact)
	http.HandleFunc("/apply", apply)

	http.ListenAndServe(":8080", nil)
}
