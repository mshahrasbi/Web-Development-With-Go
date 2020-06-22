// # Using Curl
// 1. Start your server
// ```
// go run main.go
// ```
// 1. Enter this at the terminal
// ```
// curl http://localhost:8080
// ```

package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprint(res, "Welcome!\n")
}

func main() {
	r := httprouter.New()
	r.GET("/", index)

	http.ListenAndServe("localhost:8080", r)
}
