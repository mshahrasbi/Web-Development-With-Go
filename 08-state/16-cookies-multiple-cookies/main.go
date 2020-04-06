package main

import (
	"fmt"
	"log"
	"net/http"
)

func set(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "my-cookie",
		Value: "My Value",
		Path:  "/",
	})

	fmt.Fprintln(res, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(res, " in chrome go to: dev tools / apploication / cookies")
}

func read(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(res, http.StatusText(400), http.StatusBadRequest)
		return
	} else {
		fmt.Fprintln(res, "YOUR COOKIE: ", c)
	}

	c1, err := req.Cookie("general")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(res, "YOUR COOKIE #2: ", c1)
	}

	c2, err := req.Cookie("specific")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(res, "YOUR COOKIE #3: ", c2)
	}
}

func abundance(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "general",
		Value: "Some of my general values go here",
	})

	http.SetCookie(res, &http.Cookie{
		Name:  "specific",
		Value: "Som other specific of mine values",
	})

	fmt.Fprintln(res, "COOKIES WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(res, "in chrome go to: dev tools / apploication / cookies")
}

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
