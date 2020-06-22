
package main

import (
    "log"
    "os"
    "text/template"
)

func main() {

    tpl, err := template.ParseGlob("templates/*.ghtml")
    if err != nil {
        log.Fatalln(err)
    }
    err = tpl.Execute(os.Stdout, nil)
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