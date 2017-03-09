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
  http.HandleFunc("om5z3s-b88903094z.120170301173338000g43lq21q.10.jpeg",burgerandicecoldbeer)
  http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, req *http.request) {
  tpl.ExecuteTemplate(w, "index.html", nil)
}  

func css(w http.ResponseWriter, req *http.request) {
http.ServerFile(w, req, "main.css") 
}
                 
func burgerandcoldbeer(w http.ResponseWriter, req *http.request) {
http.ServerFile(w, req, "om5z3s-b88903094z.120170301173338000g43lq21q.10.jpeg",burgerandicecoldbeer)
}
