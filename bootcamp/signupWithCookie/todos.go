package todos

import (
	"errors"
	"net/http"

	"github.com/Tomoka64/bootcamp/signupWithCookie/config"
	_ "github.com/lib/pq"
	"gopkg.in/mgo.v2/bson"
)

type Todo struct {
	Isbn  string
	Title string
	Due   string
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	tds, err := AllTodos()
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}
	config.TPL.ExecuteTemplate(w, "todos.gohtml", tds)
}

func Show(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	td, err := OneTodo(r)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	config.TPL.ExecuteTemplate(w, "show.gohtml", td)
}

func Create(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "create.gohtml", nil)
}

func CreateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	td, err := PutTodo(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}
	config.TPL.ExecuteTemplate(w, "created.gohtml", td)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	td, err := OneTodo(r)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	config.TPL.ExecuteTemplate(w, "update.gohtml", td)
}

func UpdateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	td, err := UpdateTodo(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusBadRequest)
		return
	}
	config.TPL.ExecuteTemplate(w, "updated.gohtml", td)
}

func DeleteProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	err := DeleteTodo(r)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/todos", http.StatusSeeOther)
}

func AllTodos() ([]Todo, error) {
	tds := []Todo{}
	err := config.Todos.Find(bson.M{}).All(&tds)
	if err != nil {
		return nil, err
	}
	return tds, nil
}

func OneTodo(r *http.Request) (Todo, error) {
	td := Todo{}
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return td, errors.New("400. Bad Request.")
	}
	err := config.Todos.Find(bson.M{"isbn": isbn}).One(&td)
	if err != nil {
		return td, err
	}
	return td, nil
}

func PutTodo(r *http.Request) (Todo, error) {
	td := Todo{}
	td.Isbn = r.FormValue("isbn")
	td.Title = r.FormValue("title")
	td.Due = r.FormValue("due")
	if td.Isbn == "" || td.Title == "" || td.Due == "" {
		return td, errors.New("All fields must be complete.")
	}
	err := config.Todos.Insert(td)
	if err != nil {
		return td, errors.New("Internal Error")
	}
	return td, nil
}

func UpdateTodo(r *http.Request) (Todo, error) {
	td := Todo{}
	td.Isbn = r.FormValue("isbn")
	td.Title = r.FormValue("title")
	td.Due = r.FormValue("due")
	if td.Isbn == "" || td.Title == "" || td.Due == "" {
		return td, errors.New("All fields must be complete.")
	}
	err := config.Todos.Update(bson.M{"isbn": td.Isbn}, &td)
	if err != nil {
		return td, err
	}
	return td, nil
}

func DeleteTodo(r *http.Request) error {
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return errors.New("Bad REquest")
	}
	err := config.Todos.Remove(bson.M{"isbn": isbn})
	if err != nil {
		return errors.New("INTernal S Eror")
	}
	return nil
}
