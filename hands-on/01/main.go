package main

import (
"net/http"
"html/template"
"io"
)

var tpl *template.Template

func init() {
tpl = template.must(template.ParseGlob("template/*"))
}

func main() {
  http.HandleFunc("/", foo)
  http.HandleFunc("/dog/", dog)
  http.HandleFunc("/main.css", css)
  http.HandleFunc("/dog_with_the_tv_looking_hd_picture_168923.jpeg, pic)
  http.ListenAndServe(":8080", nil)
}
                  
func foo(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, `<h1>foo ran</h1`>
}

func dog(w http.ResposeWriter, r *http.Request) {
  tpl.ExecuteTemplate(w, "dog.html", nil)
}

func css(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w,r, "main.css")
}

func pic(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w,r, "dog_with_the_tv_looking_hd_picture_168923.jpeg")
}
