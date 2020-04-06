package main

// # Step 1
// Created ```func getUser``` and put it in a new file, session.go. This refactor allows us to use the same code in index and bar.
// ```
// func getUser(w http.ResponseWriter, req *http.Request) user
// ```
//  ## IMPORTANT
//  Now that you have two go files, you cannot use "go run main.go" to get your application to run. That is only asking for one
//  go file: main.go. You need to use either "go build" and then run the executable, or "go run *.go"

//  ## WebStorm Users
//  Note to WebStorm users: when you create a new go page that has code in package main, webstorm will highlight an
//  error "multiplate packages in directory"; this will go away in time, or you can restart webstorm for it to go away immediately.

// # Step 2
// Created ```func signup``` and removed the signup code from ```func index```. A new field for password was added to the user struct.
// ```
// func signup(w http.ResponseWriter, req *http.Request)
// ```

// # Step 3
// Created ```func alreadyLoggedIn``` and put it on the session.go page. This refactor allows us to use the same code in bar and signup.
// ```
// func alreadyLoggedIn(req *http.Request) bool
// ```

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName  string
	Password  string
	FirstName string
	LastName  string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]string{} // session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func index(res http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	tpl.ExecuteTemplate(res, "index.gohtml", u)
}

func bar(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	u := getUser(req)
	tpl.ExecuteTemplate(res, "bar.gohtml", u)
}

func signup(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if req.Method == http.MethodPost {

		// get form values
		un := req.FormValue("username")
		fn := req.FormValue("firstname")
		ln := req.FormValue("lastname")
		p := req.FormValue("password")

		// username taken?

		if _, ok := dbUsers[un]; ok {
			http.Error(res, "Username already taken", http.StatusForbidden)
			return
		}

		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(res, c)
		dbSessions[c.Value] = un

		// store user in dbUsers
		u := user{un, p, fn, ln}
		dbUsers[un] = u

		// redirect
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(res, "signup.gohtml", nil)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
