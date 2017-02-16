package main
import(

"net/http"
"html/template"
)

var tpl *template.Template

func init(){
  tpl =template.Must(template.ParseGlob("templates/*"))
}
func main (){


http.HandleFunc("/",index)
http.HandleFunc("/contact",contact)
http.HandleFunc("/about",about)
http.ListenAndServe(":8080", nil)

}
func index(w http.ResponseWriter, r *http.Request){
type person struct{
First string
Last string
  }

p1 :=person{"Kenneth", "Spry"}
tpl.ExecuteTemplate(w,"index.html", p1)
}

func contact(w http.ResponseWriter, r *http.Request){
xi :=[]int{1,2,3,4,5}
tpl.ExecuteTemplate(w,"contact.html", xi)
}

func about(w http.ResponseWriter, r *http.Request){
m := map[string]int{"Michael": 23,}
tpl.ExecuteTemplate(w,"about.html", m)
}
