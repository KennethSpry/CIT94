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
func bowzer(w http.ResponseWriter, r *http.Request) {
  c := &http.Cookie{
    Name: "User-Cookie",
    Value: "Value",
    Path: "/dog/bowzer",
  }
  http.SetCookie(w, c)
  tpl.ExecuteTemplate(w, "bowzer.html", c)

}

func bowzerpics(w http.ResponseWriter, r *http.Request) {
  var c *http.Cookie
  c, err := r.Cookie("User-Cookie")
  if err != nil {
    fmt.Println(err)
  }
  tpl.ExecuteTemplate(w, "bowzerpics.html", c)


}

func cat(w http.ResponseWriter, r *http.Request) {
  var c *http.Cookie
  c, err := r.Cookie("User-Cookie")
  if err != nil {
    fmt.Println(err)
  }
  tpl.ExecuteTemplate(w, "cat.html", c)

}
