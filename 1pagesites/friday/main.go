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
  http.HandleFunc("cumulus_clouds_in_fair_weather",clouds)
  http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, req *http.request) {
  tpl.ExecuteTemplate(w, "index.html", nil)
}  

func css(w http.ResponseWriter, req *http.request) {
http.ServerFile(w, req, "main.css") 
}
                 
func imerge(w http.ResponseWriter, req *http.request) {
http.ServerFile(w, req, "Cumulus_clouds_in_fair_weather.jpeg",clouds)
}
