package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}

//
// var tmp *template.Template
//
// func main() {
// 	http.HandleFunc("/", index)
// 	http.HandleFunc("/surf.jpg", image)
// 	http.ListenAndServe(":8080", nil)
// }
//
// func index(w http.ResponseWriter, r *http.Request) {
// 	tmp, err := template.ParseFiles("index.html")
// 	if err != nil {
// 		log.Fatalln("could not parse file")
// 	}
// 	err = tmp.ExecuteTemplate(w, "index.gohtml", nil)
//
// }
//
// func image(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "surf.jpg")
// }
