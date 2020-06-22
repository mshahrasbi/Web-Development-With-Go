// # Using Curl

// 1. Start your server

// 1. Enter this at the terminal

// ```
// curl http://localhost:8080/user/1
// ```

// ```
// curl -X POST -H "Content-Type: application/json" -d '{"Name":"James Bond","Gender":"male","Age":32,"Id":"777"}' http://localhost:8080/user
// ```

// -X is short for --request
// Specifies a custom request method to use when communicating with the HTTP server.

// -H is short for --header

// -d is short for --data

// ```
// curl -X DELETE -H "Content-Type: application/json" http://localhost:8080/user/777
// ```

package main

import (
	"controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	uc := controllers.NewUserController()

	r.GET("/user/:id", uc.GetUser)

	// add another router
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe("localhost:8080", r)
}
