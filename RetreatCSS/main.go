package main

import(
  "net/http"
  "io"
)

func main() {
  http.HandleFunc("/",index)
  http.HandleFunc("/main.css",css)
  http.HandleFunc("/imerge.jpeg", imerge)
  http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, req *http.Request){
  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  io.WriteString(w, `<!DOCTYPE html>
<html lang="en">
<head>
<link rel="stylesheet" href="main.css">
</head>
<body>
<h1>Hello, Where Are You?</h1>
</body>
</html>`)
}

func css(w http.ResponseWriter, req *http.Request){
  http.ServeFile(w, req, "main.css")
}

func imerge(w http.ResponseWriter, req *http.Request){
  http.ServeFile(w, req, "imerge.jpeg")
}
