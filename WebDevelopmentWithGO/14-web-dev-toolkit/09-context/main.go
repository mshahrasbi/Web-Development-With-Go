package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func foo(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	ctx = context.WithValue(ctx, "userID", 888)
	ctx = context.WithValue(ctx, "fname", "John")

	results := dbAccess(ctx)

	fmt.Fprintln(res, results)
}

func bar(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	log.Println(ctx)
	fmt.Fprintln(res, ctx)
}

func dbAccess(ctx context.Context) int {
	uid := ctx.Value("userID").(int)
	return uid
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
