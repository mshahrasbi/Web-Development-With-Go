package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]

	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])

	str := fmt.Sprint(`
		<!DOCTYPE html>
		<html lang = "en">
			<head>
				<meta charset = "UTF-8">
				<title>Hello World!</title>
			</head>
			<h1>` + name + `</h1>
		</html>
	`)

	nf, err := os.Create("index2.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}

	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))

}

// >go run main2.go Name
