package main

import (
	"fmt"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func index(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		id, err1 := uuid.NewV4()
		if err1 != nil {
			log.Fatalln(err1)
		}
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			//Secure:   true,
			HttpOnly: true,
			Path:     "/",
		}

		http.SetCookie(res, cookie)
	}

	fmt.Println(cookie)
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
