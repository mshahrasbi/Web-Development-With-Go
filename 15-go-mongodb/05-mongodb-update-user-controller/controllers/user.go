package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mypkg/models"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

// UserController added session to our userController
type UserController struct {
	session *mgo.Session
}

// NewUserController added session to our userController
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// GetUser return users
func (us UserController) GetUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {

	u := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    32,
		Id:     p.ByName("id"),
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	res.Header().Set("COntent-Type", "application/json")
	res.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(res, "%s\n", uj)
}

// CreateUser create a user
func (us UserController) CreateUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	u := models.User{}

	json.NewDecoder(req.Body).Decode(&u)
	u.Id = "007"

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	res.Header().Set("Content-Type", "appliaction/json")
	res.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(res, "%s\n", uj)
}

// DeleteUser delete a user
func (us UserController) DeleteUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {

	// TODO: only write code to delete user
	res.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(res, "Write code to delete user\n")
}
