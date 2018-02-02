// 1. ListenAndServe on port ":8080" using the default ServeMux.
// Use HandleFunc to add the following routes to the default ServeMux:
// "/" "/dog/" "/me/
// Add a func for each of the routes.
// Have the "/me/" route print out your name.
//
// 2.Take the previous program in the previous folder and change it so that:
// a template is parsed and served
// you pass data into the template

package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

type person struct {
	Name    string
	Message string
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	// http.Handle("/", http.HandlerFunc(foo))
	// http.Handle("/dog/", http.HandlerFunc(bar))
	// http.Handle("/tm/", http.HandlerFunc(mcleod))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "doggy")
}

func me(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("tomo.gohtml")
	if err != nil {
		panic(err)
	}

	person1 := &person{
		"Tomoka",
		"hello the world",
	}

	err = tmp.ExecuteTemplate(w, "tomo.gohtml", person1)
	if err != nil {
		log.Fatalln("could not execute the file")
	}
}
