// Building upon the code from the previous problem:
// Extract the code you wrote to READ from the connection using bufio.NewScanner into its own function called "serve".
// Pass the connection of type net.Conn as an argument into this function.
// Add "go" in front of the call to "serve" to enable concurrency and multiple connections

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
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

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if ln == "" {
			// when ln is empty, header is done
			fmt.Println("THIS IS THE END OF THE HTTP REQEST HEADERS")
			break
		}
	}
}
