// # Serve the files in the "starting-files" folder

// ## To get your images to serve, use:

// ``` Go
// 	func StripPrefix(prefix string, h Handler) Handler
// 	func FileServer(root FileSystem) Handler
// ```

// Constraint: you are not allowed to change the route being used for images in the template file

package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func index(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)
}
