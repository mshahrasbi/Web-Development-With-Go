package main

// # Step 1:
// Created type session which is a struct with a lastActivity field. This will allow us to know the last time a session was used.

// # Step 2:
// Updated dbSessions to be of type map[string]session

// # Step 3:
// Updated all reads/writes to dbSessions

// # Step 4:
// Apply the MaxAge field to cookie

// # Step 5:
// Updated func alreadyLoggedIn to be able to set a cookie, adding the ResponseWriter to its parameters
// ```
// func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
// ```

// # Step 6:
// Added dbSessionsCleaned time.Time to keep track of the last time we cleaned out our sessions.

// # Step 7:
// Added func cleanSessions to remove unused sessions from dbSessions. Set it to run whenever someone logs out and a certain amount of
// time has elapsed (in production you'd set this to run during a time when the server wasn't busy).

import (
	"html/template"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName  string
	Password  []byte
	FirstName string
	LastName  string
	Role      string
}

type session struct {
	un           string
	lastActivity time.Time
}

var tpl *template.Template
var dbUsers = map[string]user{}       // user ID, user
var dbSessions = map[string]session{} // session ID, session
var dbSessionsCleaned time.Time

const sessionLength int = 30

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	dbSessionsCleaned = time.Now()
}

func index(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	showSessions()
	tpl.ExecuteTemplate(res, "index.gohtml", u)
}

func bar(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	u := getUser(res, req)
	if u.Role != "007" {
		http.Error(res, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}

	showSessions()
	tpl.ExecuteTemplate(res, "bar.gohtml", u)
}

func signup(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(res, req) {
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
		r := req.FormValue("role")

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
		c.MaxAge = sessionLength
		http.SetCookie(res, c)
		dbSessions[c.Value] = session{
			un,
			time.Now(),
		}

		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(res, "Internal server error", http.StatusInternalServerError)
			return
		}

		u := user{un, bs, fn, ln, r}
		dbUsers[un] = u

		// redirect
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(res, "signup.gohtml", nil)
}

func login(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}

	var u user
	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")

		// is there a username?
		u, ok := dbUsers[un]
		if !ok {
			http.Error(res, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(res, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = sessionLength
		http.SetCookie(res, c)
		dbSessions[c.Value] = session{
			un,
			time.Now(),
		}
		http.Redirect(res, req, "/", http.StatusSeeOther)

		return
	}

	showSessions()
	tpl.ExecuteTemplate(res, "login.gohtml", u)
}

func logout(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	c, _ := req.Cookie("session")

	// delete the session
	delete(dbSessions, c.Value)

	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}

	http.SetCookie(res, c)

	// clean up dbsessions
	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}

	http.Redirect(res, req, "/login", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
