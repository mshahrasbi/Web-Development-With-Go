package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func index(res http.ResponseWriter, req *http.Request) {
	c := getCookie(res, req)

	// process form submission
	if req.Method == http.MethodPost {
		mf, fh, err := req.FormFile("nf")
		if err != nil {
			fmt.Println(err)
		}

		defer mf.Close()

		// create sha for file name
		ext := strings.Split(fh.Filename, ".")[1]
		h := sha1.New()
		io.Copy(h, mf)
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext

		// create new file
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}

		path := filepath.Join(wd, "public", "pics", fname)
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}

		defer nf.Close()

		// copy
		mf.Seek(0, 0)
		io.Copy(nf, mf)

		// add filename to this user's cookie
		c = appendValue(res, c, fname)
	}

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

func appendValue(res http.ResponseWriter, cookie *http.Cookie, fname string) *http.Cookie {

	s := cookie.Value
	if !strings.Contains(s, fname) {
		s += "|" + fname
	}

	cookie.Value = s
	http.SetCookie(res, cookie)

	return cookie
}

func main() {
	http.HandleFunc("/", index)

	// add route to serve pictures
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
