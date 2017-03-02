package main

import(
  "net/http"
  "io"
)
func main(){
  http.HandleFunc("/", dog)
  http.HandleFunc("/snarle", snarlingdog)
  http.HandleFunc("/cupcake", cupcakeanddog)
  http.HandleFunc("/wasn't me", wasntmedog)
  http.HandleFunc("/hide", hidingdog)
  http.Handle("/pictures/", http.StripPrefix("/pictures", http.FileServer(http.Dir("./pictures"))))
  http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request){
w.Header().Set("Content-Type", "text/html; charset=utf-8")
io.WriteString(w, `<img src="/pictures/images.jpeg">`)
}

func snarlingdog(w http.ResponseWriter, req *http.Request){
w.Header().Set("Content-Type", "text/html; charset=utf-8")
io.WriteString(w, `<img src="/pictures/snarle.jpeg">`)
}

func cupcakeanddog(w http.ResponseWriter, req *http.Request){
w.Header().Set("Content-Type", "text/html; charset=utf-8")
io.WriteString(w, `<img src="/pictures/cupcake.jpeg">`)
}

func wasntmedog(w http.ResponseWriter, req *http.Request){
w.Header().Set("Content-Type", "text/html; charset=utf-8")
io.WriteString(w, `<img src="/pictures/wasn't me.jpeg">`)
}

func hidingdog(w http.ResponseWriter, req *http.Request){
w.Header().Set("Content-Type", "text/html; charset=utf-8")
io.WriteString(w, `<img src="/pictures/hide.jpeg">`)
}
