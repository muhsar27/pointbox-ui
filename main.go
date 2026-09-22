package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type Member struct {
	Name        string
	Points      int
	LastUpdated string
}

var members = make(map[int]Member)

func main() {

	members[1] = Member{"Muhammad", 1500, time.Now().Format("02-01-2006 03:04PM")}
	members[2] = Member{"Ebuka", 1500, time.Now().Format("02-01-2006 03:04PM")}


	mux := http.NewServeMux()

	mux.HandleFunc("GET /", home)

	log.Println("Server starting in 5...4...3...2...1")

	err := http.ListenAndServe(":5000", mux)
	log.Fatal(err)

}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./home.html")

	if err != nil {
		log.Print(err.Error())
		return
	}

	err = tmpl.Execute(w, members)
	if err != nil {
		log.Print(err.Error())
		return
	}
	//show members

}
