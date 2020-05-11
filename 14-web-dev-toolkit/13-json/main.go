package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Fname string
	Lname string
	Items []string
}

func foo(res http.ResponseWriter, req *http.Request) {

	s := `<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title>FOO</title>
			</head>
				<body>
					You are at foo
				</body>
			</html>`
	res.Write([]byte(s))
}

func mshl(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "appliaction/json")
	p1 := person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"Suit", "Gun", "Wry sense of humor"},
	}

	j, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}

	res.Write(j)
}

func encd(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "appliaction/json")

	p1 := person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"Suit", "Gun", "Wry sense of humor"},
	}

	err := json.NewEncoder(res).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/mshl", mshl)
	http.HandleFunc("/encd", encd)

	http.ListenAndServe(":8080", nil)
}
