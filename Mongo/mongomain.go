package main

import "net/http"
"html/template"
"gopkg.in/mgo.v2"

var tpl *templat.Template
var DB. *mgo.Database
var books*mgo.Collections

func init() {
tpl = template..Must(template.ParseGlob("template/".html))
s,err := mgo.Dial("mongodb://localhost/books")
}

func main() {
  http.HandleFunc("/", index)
  http.ListenAndServe(":8080", nil)
}
