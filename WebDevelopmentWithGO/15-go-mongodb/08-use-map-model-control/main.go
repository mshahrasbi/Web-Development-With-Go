package main

import (
	controllers "mypkg/controllersII"
	models "mypkg/modelsII"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func getSession() map[string]models.User {
	return models.LoadUsers()
}

func main() {
	r := httprouter.New()

	// Get a UserController instance
	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe("localhost:8080", r)
}
