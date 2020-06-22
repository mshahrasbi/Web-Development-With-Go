// Building upon the code from the previous problem:

// Change your RESPONSE HEADER "content-type" from "text/plain" to "text/html"
// Change the RESPONSE from "CHECK OUT THE RESPONSE BODY PAYLOAD" (and everything else it contained: request method,
// request URI) to an HTML PAGE that prints "HOLY COW THIS IS LOW LEVEL" in <h1> tags.

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go serve(conn)

	}

}

func serve(c net.Conn) {

	defer c.Close()
	scanner := bufio.NewScanner(c)

	var i int
	var rMethod, rURI string

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			// we are in REQUEST LINE
			xs := strings.Fields(ln)
			if len(xs) > 1 {
				rMethod = xs[0]
				rURI = xs[1]

				fmt.Println("METHOD:", rMethod)
				fmt.Println("URI:", rURI)
			}
		}

		if ln == "" {
			// when ln is empty, header is done
			fmt.Println("THIS IS THE END OF THE HTTP REQEST HEADERS")
			break
		}
	}

	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Code Gangsta</title>
		</head>
		<body>
			<h1>HOLY COW THIS IS LOW LEVEL</h1>
		</body>
		</html>
	`

	io.WriteString(c, "HTTP/1.1 200 OK\r\n")

	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(c, "Content-Type: text/html\r\n")

	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}
