package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	// tpl = template.Must(template.ParseFiles("tpl.gohtml"))
	tpl = template.Must(template.ParseFiles("tpl1.gohtml"))
}

func main() {
	sages := []string{"Gandi", "MLK", "Buddha", "Jesus", "Mohammad"}

	err := tpl.Execute(os.Stdout, sages)

	if err != nil {
		log.Fatalln(err)
	}
}
