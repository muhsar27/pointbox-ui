package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Member struct {
	Name        string
	Points      int
	LastUpdated string
}

var members = make(map[int]Member)
var currentId int = 1

func main() {

	//members[1] = Member{"Muhammad", 1500, time.Now().Format("02-01-2006 03:04PM")}
	//members[2] = Member{"Ebuka", 1500, time.Now().Format("02-01-2006 03:04PM")}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", home)
	mux.HandleFunc("GET /add", addMember)
	mux.HandleFunc("POST /add", addMemberPost)
	mux.HandleFunc("POST /add/points/{ID}", addPoint)
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	log.Println("Server starting in http://localhost:5000")

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

func addMember(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./addmember.html")

	if err != nil {
		log.Print(err.Error())
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		return
	}
}

func addMemberPost(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	point, err := strconv.Atoi(r.FormValue("point"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	members[currentId] = Member{name, point, time.Now().Format("02-01-2006 03:04PM")}
	currentId++

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func addPoint(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("ID"))

	if err != nil {
		http.NotFound(w, r)
		return
	}

	_, ok := members[id]

	if !ok {
		http.NotFound(w, r)
		return
	}

	user := members[id]
	user.Points += 10
	user.LastUpdated = time.Now().Format("02-01-2006 03:04PM")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
