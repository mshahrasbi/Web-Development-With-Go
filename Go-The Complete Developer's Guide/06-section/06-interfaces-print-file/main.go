package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// fmt.Println(os.Args[1])
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
