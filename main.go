package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/morelmiles/zuri-internship/luigi"
	gomail "gopkg.in/mail.v2"
)

var tpl *template.Template

type Content struct {
		FirstName string
		SecondName string
		Email string
		Message	string
		Reason string
	}

func init () {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml" ))
}
func main() {
	//Task One: Print name to the console 
	luigi.LuigiMorel()


// Task 2: Resumel=
r := mux.NewRouter()
r.HandleFunc("/", index)
r.HandleFunc("/process", process)}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func send(w http.ResponseWriter, r *http.Request) {
	 m := &Content{
	 Email: r.FormValue("emailAddress"),
	FirstName : r.FormValue("firstName"),
	SecondName : r.FormValue("secondName"),
	Reason: r.FormValue("reason"),
	Message: r.FormValue("message"),
	}

	err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
  http.Redirect(w, r, "/process", http.StatusSeeOther)

}

func process(w http.ResponseWriter, r *http.Request) {
    tpl.ExecuteTemplate(w, "templates/process.gohtml", nil)
}


func (m *Content ) Deliver() {
	m := gomail.NewMessage()
	m.SetHeader("From", "Luigi Morel <morelmiles@gmail.com>")
    m.SetHeader("To", "Luigi Morel <jsmith@gmail.com>")
    m.SetAddressHeader("reply-to", "m.Email")
    m.SetHeader("Subject", "Contact")

    m.SetBody("text/html", fmt.Sprintf("<b>First Name</b>: %s", m.FirstName))
	m.SetBody("text/html", fmt.Sprintf("<b>Second Name</b>: %s", m.SecondName))
    m.SetBody("text/html", fmt.Sprintf("<b>Email</b>: %s", m.Email))
	m.SetBody("text/html", fmt.Sprintf("<b>Reason</b>: %s", m.Reason))
    m.SetBody("text/html", fmt.Sprintf("<b>Message</b>: %s", m.Message))

    deliver := gomail.NewDialer("smtp.gmail.com", 587, "morelmiles@gmail.com", "password")
    if err := d.DialAndSend(m); err != nil {
        panic(err)
    }
}