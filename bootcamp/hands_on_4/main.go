package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/gopher.png", dogpic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src = "gopher.png">`)
}

func dogpic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "gopher.png")
}
