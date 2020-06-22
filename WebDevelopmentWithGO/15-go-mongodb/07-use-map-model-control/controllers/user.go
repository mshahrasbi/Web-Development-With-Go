package controllers

import (
	"encoding/json"
	"fmt"
	models "mypkg/modelsI"
	"net/http"

	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
)

// UserController user controller
type UserController struct {
	session map[string]models.User
}

// NewUserController new User Controller
func NewUserController(m map[string]models.User) *UserController {
	return &UserController{m}
}

// GetUser Get the User
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// grap id
	id := p.ByName("ID")

	// retrieve user
	u := uc.session[id]

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200

	fmt.Fprintf(w, "%s\n", uj)
}

// CreateUser create a user
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	// create ID
	uID, err := uuid.NewV4()
	u.ID = uID.String()

	// store the user
	uc.session[u.ID] = u

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

// DeleteUser delete user
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	delete(uc.session, id)

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", id, "\n")
}
