// # Use a map
// Instead of using mongodb, store all of the data in a map.

package main

import (
	controllers "mypkg/controllersI"
	models "mypkg/modelsI"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func getSession() map[string]models.User {
	return make(map[string]models.User)
}

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:ID", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}
