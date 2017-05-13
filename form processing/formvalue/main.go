package main

import (
  "net/http"
  "html/template"
  "log"
)

var tpl *template.Template

func init() {
  tpl = template.Must(template.ParseFiles("templates/index.html"))
}

type order struct {
  Name string
  Email string
  Phone string
  Quantity string
}

func main() {
  http.HandleFunc("/", index)
  http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

  n := req.FormValue("name")
  e := req.FormValue("email")
  p := req.FormValue("phone")
  q := req.FormValue("quantity")

  err := tpl.ExecuteTemplate(w, "index.html", order{n,e,p,q})
  if err !=nil {
    http.Error(w, err.Error(), 500)
    log.Fatalln(err)
  }
}

    
