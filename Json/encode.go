package main

import (
  "net/http"
  "encoding/json"
  "fmt"
  "html/template"
)

type player struct{
  FirstName string
  LastName string
  UserName string
  Age int
  PremiumMember bool
}

var tpl *template.Template
var xpl []player

func init()  {
  tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
  
  http.HandleFunc("/", index)
  http.HandleFunc("/send", send)
  http.HandleFunc("/catch", catch)
    http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
   p1 := player{
    FirstName: "Kenneth",
    LastName: "Spry",
    UserName: "Shadowhawk",
    Age: 55,
    PremiumMember: false,
  }
  p2 := player{
    FirstName: "Micheal",
    LastName: "Robby",
    UserName: "Elite H@xor",
    Age: 23,
    PremiumMember: true,
  }
  
    xpl = []player{p1,p2}
  err := json.NewEncoder(w).Encode(xpl)
  if err != nil {
    fmt.Println(err)
  }
 
}

func send(w http.ResponseWriter, r *http.Request) {
   tpl.ExecuteTemplate(w, "index.html", nil)
}


func catch(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST"{

    err := json.NewDecoder(r.Body).Decode(&xpl)
    if err != nil {
      fmt.Println(err)
      return
    }
}

  fmt.Println(xpl)
  tpl.ExecuteTemplate(w, "catch.html", xpl)
}
