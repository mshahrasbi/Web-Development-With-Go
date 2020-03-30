package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	// tpl = template.Must(template.ParseFiles("tpl.gohtml"))
	tpl = template.Must(template.ParseFiles("tpl1.gohtml"))
}

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	mlk := sage{
		Name:  "MLK",
		Motto: "Hatred never ceases with hatred but with love alone is healed",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}

	mohammad := sage{
		Name:  "Mohammad",
		Motto: "To overcome evil with good is good, to resist evil by evil is evil",
	}

	sages := []sage{buddha, gandhi, mlk, jesus, mohammad}

	err := tpl.Execute(os.Stdout, sages)

	if err != nil {
		log.Fatalln(err)
	}
}
