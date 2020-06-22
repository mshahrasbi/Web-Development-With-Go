package main

import (
	"fmt"
	"time"
)

// gen is a broken generator that will leak a goroutine
func gen() <-chan int {
	channel := make(chan int)
	go func() {
		var n int
		for {
			channel <- n
			n++
		}
	}()

	return channel
}

func main() {
	for n := range gen() {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}

	time.Sleep(1 * time.Minute)
}
