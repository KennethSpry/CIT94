package main

import (
  "net/http"
  "html/template"
  "fmt"
)

var tpl *template.Template

func init() {
  tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
  http.HandleFunc("/", index)
  http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
  var c *http.Cookie
  c, err := r.Cookie("user-Cookie")
  if err |= nil {
    fmt.Println(err)
  }
  tpl.ExecuteTemplate(w, "index.html", c)
}
