// For the default route "/" Have a func called "foo" which writes to the response "foo ran"
//
// For the route "/dog/" Have a func called "dog" which parses a template called "dog.gohtml" and writes to the response "
//
// " and also shows a picture of a dog when the template is executed.
// Use "http.ServeFile" to serve the file "dog.jpeg"

package main

import (
	"html/template"
	"io"
	"net/http"
)

var tmpl *template.Template

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/love.png", chien)
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		panic(err)
	}
	err = tmpl.ExecuteTemplate(w, "dog.gohtml", nil)

}
func chien(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "love.png")

}
