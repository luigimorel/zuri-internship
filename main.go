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
		reason := r.FormValue("reason")
		message := r.FormValue("message")


		type Content struct {
		FirstName string
		SecondName string
		Email string
		Message	string
		Reason string
		} {
		FirstName: firstName, 
		SecondName: secondName,
		Email: emailAddress,
		Reason: reason,
		Message: message,
		}
			tpl.ExecuteTemplate(w, "process.gohtml", values)
		}

func (m *Content ) Deliver() {
	m := gomain.NewMessage()
	 m.SetHeader("From", "Luigi Morel <morelmiles@gmail.com>")
    m.SetHeader("To", "Luigi Morel <jsmith@gmail.com>")
    m.SetAddressHeader("reply-to", "m.Email")
    m.SetHeader("Subject", "Contact")
    m.SetBody("text/html", "<b>First Name</b>: m.FirstName")
	m.SetBody("text/html", "<b>Second Name</b>: m.SecondName")
    m.SetBody("text/html", "<b>Email</b>: m.Email")
    m.SetBody("text/html", "<b>Message</b>: m.Message")
	m.SetBody("text/html", "<b>Reason</b>: m.Reason")

    deliver := gomail.NewDialer("smtp.gmail.com", 587, "morelmiles@gmail.com", "password")
    if err := d.DialAndSend(m); err != nil {
        panic(err)
    }
}