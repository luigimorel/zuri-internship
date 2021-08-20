package main

import (
	"html/template"
	"net/http"

	"github.com/morelmiles/zuri-internship/luigi"
)

var tpl *template.Template

func init () {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml" ))
}
func main() {
	//Task One: Print name to the console 
	luigi.LuigiMorel()


// Task 2: Resume 
http.HandleFunc("/", index)
http.HandleFunc("/process", processor)
http.ListenAndServe(":8080", nil)	
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	return 
}	

	firstName := r.FormValue("firstName")
		secondName := r.FormValue("secondName")
		emailAddress := r.FormValue("emailAddress")
		message := r.FormValue("message")

values := struct {
FirstName string
SecondName string
Email string
Message	string
} {
FirstName: firstName, 
SecondName: secondName,
Email: emailAddress,
Message: message,
}
	tpl.ExecuteTemplate(w, "process.gohtml", values)
}
