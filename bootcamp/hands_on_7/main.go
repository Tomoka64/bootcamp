package main

// //
// IN A GO FILE
//
// Create a variable with the identifier "rcvd" of type string.
//
// Store a raw string literal of the JSON created in the previous step as the value of the variable "rcvd".
//
// Unmarshal "rcvd" into a data structure with the identifier "data"
//
// Use a for range loop to iterate through "data" displaying the results to the terminal
import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type code struct {
	Code    int
	Descrip string
}

var tmp *template.Template

func main() {
	http.HandleFunc("/", display)
	http.ListenAndServe(":8080", nil)
}

func init() {
	tmp = template.Must(template.ParseFiles("index.html"))

}
func display(w http.ResponseWriter, r *http.Request) {
	var data []code
	rcvd := `[{"Code":200,"Descrip":"StatusOK"},{"Code":301,"Descrip":"StatusMovedPermanently"},{"Code":302,"Descrip":"StatusFound"},{"Code":303,"Descrip":"StatusSeeOther"},{"Code":307,"Descrip":"StatusTemporaryRedirect"},{"Code":400,"Descrip":"StatusBadRequest"},{"Code":401,"Descrip":"StatusUnauthorized"},{"Code":402,"Descrip":"StatusPaymentRequired"},{"Code":403,"Descrip":"StatusForbidden"},{"Code":404,"Descrip":"StatusNotFound"},{"Code":405,"Descrip":"StatusMethodNotAllowed"},{"Code":418,"Descrip":"StatusTeapot"},{"Code":500,"Descrip":"StatusInternalServerError"}]`

	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		panic(err)
	}
	for _, v := range data {
		fmt.Println(v.Code, "-", v.Descrip)
	}

	err = tmp.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		panic(err)
	}

}
