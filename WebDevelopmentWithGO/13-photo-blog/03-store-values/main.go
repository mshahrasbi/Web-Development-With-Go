package main

import (
	"html/template"
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func index(res http.ResponseWriter, req *http.Request) {
	c := getCookie(res, req)
	c = appendValue(res, c)

	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(res, "index.gohtml", xs)
}

// function to get cookie
func getCookie(res http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

		http.SetCookie(res, c)
	}

	return c
}

func appendValue(res http.ResponseWriter, cookie *http.Cookie) *http.Cookie {
	// values
	p1 := "disneyland.jpg"
	p2 := "atbeach.jpg"
	p3 := "hollywood.jpg"

	// append
	s := cookie.Value
	if !strings.Contains(s, p1) {
		s += "|" + p1
	}
	if !strings.Contains(s, p2) {
		s += "|" + p2
	}
	if !strings.Contains(s, p3) {
		s += "|" + p3
	}

	cookie.Value = s
	http.SetCookie(res, cookie)

	return cookie
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
