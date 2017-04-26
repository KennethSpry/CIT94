package main

import (
	"github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

type user struct {
	Username string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		c = &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, c)
	}
	var u user
	if uid, ok := dbSessions[c.Value]; ok {
		u = dbUsers[uid]
	}
	// Process Form Submission
	if r.Method == "POST" {
		un := r.FormValue("username")
		f := r.FormValue("first")
		l := r.FormValue("last")
		u := user{un, f, l}
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}

	tpl.ExecuteTemplate(w, "index.html", u)
}
