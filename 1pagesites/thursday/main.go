package main

import(
  "net/http"
  "io"
)

func main() {
  http.HandleFunc("/",index)
  http.HandleFunc("main.css,css)
  http.HandleFunc("imerge.jpeg",imerge)
  http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, req *http.request) {
w.Header().Set("Content-Type": "text/html; charset=utf-8")
io.WriteString(w, "template") 
}  

func css(
  
  
  func template(
    
    
    func imerge(
