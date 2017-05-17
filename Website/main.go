package main

import (
  "github.com/satori/go.uuid"
  "golang.org/x/crypto/bcrypt"
  "html/template"
  "net/http"
  "time"
)

type user struct {
  UserName string
  Password []byte
  First    string
  Last     string
  Role     string
}

type session struct {
  un           string
  lastActivity time.Time
}

var tpl *template.Template
var dbUsers = map[string]user{}      
var dbSessions = map[string]session{} 
var dbSessionsCleaned time.Time

const sessionLength int = 30

func init() {
  tpl = template.Must(template.ParseGlob("templates/*"))
  dbSessionsCleaned = time.Now()
}

func main() {
  http.HandleFunc("/", index)
  http.HandleFunc("/myplace",myplace)
  http.HandleFunc("/signup", signup)
  http.HandleFunc("/login", login)
  http.HandleFunc("/logout", logout)
  http.Handle("/favicon.ico", http.NotFoundHandler())
  http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
  u := getUser(w, req)
  showSessions() 
  tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
  u := getUser(w, req)
  if !alreadyLoggedIn(w, req) {
    http.Redirect(w, req, "/", http.StatusSeeOther)
    return
  }
  if u.Role != "021" {
    http.Error(w, "You must be 021 to enter this site", http.StatusForbidden)
    return
  }
  showSessions() 
  tpl.ExecuteTemplate(w, "myplace.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
  if alreadyLoggedIn(w, req) {
    http.Redirect(w, req, "/", http.StatusSeeOther)
    return
  }
  var u user
   if req.Method == http.MethodPost {
       un := req.FormValue("username")
    p := req.FormValue("password")
    f := req.FormValue("firstname")
    l := req.FormValue("lastname")
    r := req.FormValue("role")
        if _, ok := dbUsers[un]; ok {
      http.Error(w, "Username already taken", http.StatusForbidden)
      return
    }
        sID := uuid.NewV4()
    c := &http.Cookie{
      Name:  "session",
      Value: sID.String(),
    }
    c.MaxAge = sessionLength
    http.SetCookie(w, c)
    dbSessions[c.Value] = session{un, time.Now()}
       bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
    if err != nil {
      http.Error(w, "Internal server error", http.StatusInternalServerError)
      return
    }
    u = user{un, bs, f, l, r}
    dbUsers[un] = u
       http.Redirect(w, req, "/", http.StatusSeeOther)
    return
  }
  showSessions() 
  tpl.ExecuteTemplate(w, "signup.gohtml", u)
}

func login(w http.ResponseWriter, req *http.Request) {
  if alreadyLoggedIn(w, req) {
    http.Redirect(w, req, "/", http.StatusSeeOther)
    return
  }
  var u user
    if req.Method == http.MethodPost {
    un := req.FormValue("username")
    p := req.FormValue("password")
      u, ok := dbUsers[un]
    if !ok {
      http.Error(w, "Username and/or password do not match", http.StatusForbidden)
      return
    }
   
    err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
    if err != nil {
      http.Error(w, "Username and/or password do not match", http.StatusForbidden)
      return
    }
  
    sID := uuid.NewV4()
    c := &http.Cookie{
      Name:  "session",
      Value: sID.String(),
    }
    c.MaxAge = sessionLength
    http.SetCookie(w, c)
    dbSessions[c.Value] = session{un, time.Now()}
    http.Redirect(w, req, "/", http.StatusSeeOther)
    return
  }
  showSessions() 
  tpl.ExecuteTemplate(w, "login.gohtml", u)
}

func logout(w http.ResponseWriter, req *http.Request) {
  if !alreadyLoggedIn(w, req) {
    http.Redirect(w, req, "/", http.StatusSeeOther)
    return
  }
  c, _ := req.Cookie("session")
    delete(dbSessions, c.Value)
    c = &http.Cookie{
    Name:   "session",
    Value:  "",
    MaxAge: -1,
  }
  http.SetCookie(w, c)

 
  if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
    go cleanSessions()
  }

  http.Redirect(w, req, "/login", http.StatusSeeOther)
}

    
