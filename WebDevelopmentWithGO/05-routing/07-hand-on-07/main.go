// Building upon the code from the previous problem:
// Print to standard out (the terminal) the REQUEST method and the REQUEST URI from the REQUEST LINE.
// Add this data to your REPONSE so that this data is displayed in the browser.

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

	body := "CHECK OUT THE RESPONSE BODY PAYLOAD"
	body += "\nMethod: "
	body += rMethod
	body += "\nURI: "
	body += rURI

	io.WriteString(c, "HTTP/1.1 200 OK\r\n")

	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(c, "Content-Type: text/plain\r\n")

	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}
