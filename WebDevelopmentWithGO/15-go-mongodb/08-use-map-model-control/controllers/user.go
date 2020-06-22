package controllers

import (
	"encoding/json"
	"fmt"
	models "mypkg/modelsII"
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/julienschmidt/httprouter"
)

// UserController UserController
type UserController struct {
	session map[string]models.User
}

// NewUserController NewUserController
func NewUserController(m map[string]models.User) *UserController {
	return &UserController{m}
}

// GetUser GetUser
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Retrieve user
	u := uc.session[id]

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

// CreateUser  CreateUser
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	// create ID
	uID, err := uuid.NewV4()
	u.ID = uID.String()

	// store the user
	uc.session[u.ID] = u
	models.StoreUsers(uc.session)

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

// DeleteUser DeleteUser
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	delete(uc.session, id)

	models.StoreUsers(uc.session)

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user ", id, "\n")
}
