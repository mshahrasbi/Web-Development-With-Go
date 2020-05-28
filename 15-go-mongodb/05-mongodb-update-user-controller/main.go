// # Install Mongo
// # Go get driver for mongo
// ```
// go get gopkg.in/mgo.v2
// go get gopkg.in/mgo.v2/bson
// ```
// # In this step:
// Don't run this code
// Just making updates - a several step process.
// We will need a mongo session to use in the CRUD methods.
// We need our UserController to have access to a mongo session.
// Let's add this to controllers/user.go
// ```
// UserController struct {
//     session *mgo.Session
// }
// ```
// And now add this to controllers/user.go
// ```
// func NewUserController(s *mgo.Session) *UserController {
//     return &UserController{s}
// }
// ```
// And now add this to main.go
// ```
// func getSession() *mgo.Session {
// 	// Connect to our local mongo
// 	s, err := mgo.Dial("mongodb://localhost")

// 	// Check if connection error, is mongo running?
// 	if err != nil {
// 		panic(err)
// 	}
// 	return s
// }
// ```
// and this
// ```
// uc := controllers.NewUserController(getSession())
// ```
// 1. Enter this at the terminal
// ```
// curl http://localhost:8080/user/1
// ```

package main

import (
	"net/http"

	"github.com/mypkg/controllers-I"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func getSession() *mgo.Session {

	// connect to our mongo
	s, err := mgo.Dial("mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass%20Community&ssl=false")

	// check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}

func main() {

	r := httprouter.New()

	// get a userCOntroller instance
	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe("localhost:8080", r)
}
