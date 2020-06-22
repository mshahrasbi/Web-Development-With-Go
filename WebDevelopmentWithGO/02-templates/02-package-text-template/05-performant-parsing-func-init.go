
package main

import (
    "log"
    "os"
    "text/template"
)

var tpl *template.Template


func init() {
    tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

    err := tpl.Execute(os.Stdout, nil)
    if err != nil {
        log.Fatalln(err)
    }

    err = tpl.ExecuteTemplate(os.Stdout, "vespa.ghtml", nil)
    if err != nil {
        log.Fatalln(err)
    }

    err = tpl.ExecuteTemplate(os.Stdout, "two.ghtml", nil)
    if err != nil {
        log.Fatalln(err)
    }

    err = tpl.ExecuteTemplate(os.Stdout, "one.ghtml", nil)
    if err != nil {
        log.Fatalln(err)
    }
}