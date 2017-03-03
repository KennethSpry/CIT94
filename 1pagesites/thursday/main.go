package main

import(
  "net/http"
  "io"
)

func main() {
  http.HandleFunc("/",index)
  http.HandleFunc("/main.css,css)
  http.HandleFunc("small-island.jpeg",island)
  http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, req *http.request) {
w.Header().Set("Content-Type": "text/html; charset=utf-8")
io.WriteString(w, req "template") 
}  

func css(w http.ResponseWriter, req *http.request) {
http.ServerFile(w, req, "main.css") 
}
                  
func template(w http.ResponseWriter, req *http.request) {
http.ServerFile(w, req, "template")
}
                  
func imerge(w http.ResponseWriter, req *http.request) {
http.ServerFile(w, req, `<img scr="small-island.jpeg">`)
}
