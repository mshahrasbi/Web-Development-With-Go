package main

import "fmt"

func main() {
	name := "Jim Smith"

	tpl := `
		<!DOCTYPE html>
		<html lang = "en">
			<head>
				<meta charset = "UTF-8">
				<title>Hello World!</title>
			</head>
			<h1>` + name + `</h1>
		</html>
	`

	fmt.Println(tpl)
}

// >go run main.go
// >go run main.go > index.html
