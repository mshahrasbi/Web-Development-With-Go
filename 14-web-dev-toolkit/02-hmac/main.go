package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func index(res http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("session")
	if err != nil {
		cookie = &http.Cookie{
			Name:  "session",
			Value: "",
		}
	}

	if req.Method == http.MethodPost {
		e := req.FormValue("email")
		cookie.Value = e + `|` + getCode(e)
	}

	http.SetCookie(res, cookie)

	io.WriteString(res, `
		<!DOCTYPE html>
		<html>
			<body>
				<form method="POST">
					<input type="email" name="email">
					<input type="submit">
				</form>
				<a href="/authenticate">Validate This `+cookie.Value+`</a>
			</body>
		</html>
	`)
}

func auth(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if cookie.Value == "" {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	xs := strings.Split(cookie.Value, "|")
	email := xs[0]
	codeRcvd := xs[1]
	codeCheck := getCode(email)

	if codeRcvd != codeCheck {
		fmt.Println("HMAC code didn't match")
		fmt.Println(codeRcvd)
		fmt.Println(codeCheck)

		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	io.WriteString(res, `
		<!DOCTYPE html>
		<html>
			<body>
				<h1>`+codeRcvd+` - RECEIVED </h1>
				<h1>`+codeCheck+` - RECALCULATED </h1>
			</body>
		</html>
	`)
}

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("OurKey"))
	io.WriteString(h, data)

	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/authenticate", auth)
	http.ListenAndServe(":8080", nil)
}
