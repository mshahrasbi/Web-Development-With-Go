package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println(resp)
	fmt.Println("---------------------------------------------------------------------------------")
	fmt.Println(resp.Header)
	fmt.Println("---------------------------------------------------------------------------------")

	// make function is a built in function in the language that takes a type of a slice and as a second
	// argument this is the number of elements or empty spaces thatwe want to slice to be initialize with
	// a byte slice cangrow and shrink, but if we want to we can create a byte slice with n number of
	// empty elements inside of it. this will give us a empty byte slice that has space for 99999 elements
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
