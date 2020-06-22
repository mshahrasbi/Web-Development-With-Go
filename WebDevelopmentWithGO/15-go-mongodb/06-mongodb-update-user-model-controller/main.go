// MongoDB represents JSON documents in binary-encoded format called BSON behind the scenes. BSON extends the JSON model to provide
// additional data types and to be efficient for encoding and decoding within different languages.

// We will update our user model to change the type of our Id field to be a bson.ObjectId

// Add this to models/user.go

// ```
// type User struct {
// 	Name   string        `json:"name" bson:"name"`
// 	Gender string        `json:"gender" bson:"gender"`
// 	Age    int           `json:"age" bson:"age"`
// 	Id     bson.ObjectId `json:"id" bson:"_id"`
// }

// ```

package main

import (
	"mypkg/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")

	if err != nil {
		panic(err)
	}
	return s
}

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe("localhost:8080", r)
}
