package main

import (
	"fmt"
	"html/template"
	"log"

	//"time"

	//"html/template"
	//"log"
	//"Datarec/Sample"
	"net/http"

	uuid "github.com/satori/go.uuid"
	//"github.com/go-playground/validator"
	//"strings"
)

// func sayhelloName(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
// 	// attention: If you do not call ParseForm method, the following data can not be obtained form
// 	fmt.Println(r.Form) // print information on server side.
// 	fmt.Println("path", r.URL.Path)
// 	fmt.Println("scheme", r.URL.Scheme)
// 	fmt.Println(r.Form["url_long"])
// 	for k, v := range r.Form {
// 		fmt.Println("key:", k)
// 		fmt.Println("val:", strings.Join(v, ""))
// 	}
// 	fmt.Fprintf(w, "Hello astaxie!") // write data to response
// }

type userDetails struct {
	userName string
	password string
}

type enteredDetails struct {
	userName string
	password string
}

type notValid struct {
	Not  string
	Name string
}

var session = map[string]string{}

var details enteredDetails

var user userDetails

var stat bool

func login(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method) //get request methe

	if r.Method == "GET" {

		if _, ok := session[details.userName]; ok {
			http.Redirect(w, r, "/home", http.StatusSeeOther)

		} else {

			t, _ := template.ParseFiles("index.html")
			t.Execute(w, nil)

		}

	} else if r.Method == "POST" {

		user.userName = "Ashiq@123"
		user.password = "123456"

		details.userName = r.FormValue("username")
		details.password = r.FormValue("password")

		if details.userName == user.userName && details.password == user.password {

			stat = true
			cookie, err := r.Cookie("cook")

			if err != nil {
				id := uuid.NewV4()
				cookie = &http.Cookie{
					Name:     "cook1",
					Value:    id.String(),
					HttpOnly: true,
				}
				http.SetCookie(w, cookie)
			}

			session[details.userName] = cookie.Value

			fmt.Println(cookie)

			http.Redirect(w, r, "/home", http.StatusSeeOther)
		} else {
			t, _ := template.ParseFiles("index.html")
			p := notValid{Not: "Invalid username or password"}
			t.Execute(w, p)
		}

	}

}

func back(w http.ResponseWriter, r *http.Request) {
	if stat {
		delete(session, details.userName)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("home.html")
	h := notValid{Name: details.userName}
	t.Execute(w, h)

}

func errlo(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")

	t.Execute(w, nil)
}

func main() {
	//http.HandleFunc("/", sayhelloName) // setting router rule
	http.HandleFunc("/", login)
	http.HandleFunc("/home", home)
	http.HandleFunc("/error", errlo)
	http.HandleFunc("/back", back)

	err := http.ListenAndServe(":9890", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)

	}

}
