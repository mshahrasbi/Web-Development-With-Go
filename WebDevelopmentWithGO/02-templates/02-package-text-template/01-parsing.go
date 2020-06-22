
package main

import (
    "log"
    "os"
    "text/template"
)


func main() {

    tpl, err := template.ParseFiles("tpl.gohtml")
    if err != nil {
        log.Fatalln(err)
    }

    err = tpl.Execute(os.Stdout, nil)
    if err != nil {
        log.Fatalln(err)
    }
}

// to run it:
// >go run 01-parsing.go