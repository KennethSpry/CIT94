package main

import(
  "net/http"
  "html/template"
)
var tpl *template.Template

func init() {
  tpl =template.Must(template.ParseGlob("templates/*"))
  
func main() {
  http.HandleFunc("/",index)
  http.HandleFunc("/main.css",css)
  http.HandleFunc("mountains.jpeg",escape)
  http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, req *http.request) {
  tpl.ExecuteTemplate(w, "index.html", nil)
}  

func css(w http.ResponseWriter, req *http.request) {
http.ServerFile(w, req, "main.css") 
}
                 
func escape(w http.ResponseWriter, req *http.request) {
http.ServerFile(w, req, "mountains.jpeg",escape)
}
