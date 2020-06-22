package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func foo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	ctx = context.WithValue(ctx, "userID", 888)
	ctx = context.WithValue(ctx, "fname", "John")

	results, err := dbAccess(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}

	fmt.Fprintln(w, results)
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}

func dbAccess(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	channel := make(chan int)

	go func() {
		uid := ctx.Value("userID").(int)

		time.Sleep(10 * time.Second)

		// check to make sure we are not running in Vain
		// if ctx.Done() has
		if ctx.Err() != nil {
			log.Println("error is not nil")
			return
		}

		log.Println(uid)
		channel <- uid
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case i := <-channel:
		return i, nil
	}
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
