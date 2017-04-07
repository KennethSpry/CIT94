package main

import (
  "net/http"
  "html/template"
)

var tpl *template.Template

func init()  {
  tpl = template.Must(template.ParseFiles("templates/index.html"))
}

func main() {
  fs := http.FileServer(http.Dir("public pics"))
  http.Handle("/public pics/", fs)
  http.HandleFunc("/", index)
  http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
  tpl.ExecuteTemplate(w, "index.html", nil)
}
