package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Jim Smith"

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

	nf, err := os.Create("index1.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}

	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))

}

// >go run main1.go
