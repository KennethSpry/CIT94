package main

import (
  "io"
  "net/http"
)
func main() {
  http.HandleFunc("/", tiger)
  http.HandleFunc("/tiger_1.jpeg", tigerpic)
  http.ListenAndServe(":8080", nil)
}

func tiger(w http.ResponseWriter, req *http.Request){
  w.Header().Set("Content-Type" ,"text/html; charset=utf-8")
  io.WriteString(w, `<img scr="tiger_1.jpeg"> `)
}

func tigerPic(w http.ResponseWriter, req *http.Request) {
  http.ServerFile(w, req, "serverfile/img/tiger_1.jpeg")
}
