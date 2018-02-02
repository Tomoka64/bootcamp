// Create a data structure to pass to a template which
// 1 contains information about California hotels including Name, Address, City, Zip, Region
// 2 region can be: Southern, Central, Northern
// 3 can hold an unlimited number of hotels
package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
}

type region struct {
	Region string
	Hotels []hotel
}

type Region []region

var tmpl *template.Template

func main() {
	h := region{
		Region: "southern",
		Hotels: []hotel{
			hotel{
				Name:    "Hotel California",
				Address: "42 Sunset Boulevard",
				City:    "Los Angeles",
				Zip:     "95612",
			},
			hotel{
				Name:    "H",
				Address: "4",
				City:    "L",
				Zip:     "95612",
			},
		},
	}
	tmpl, _ = template.ParseFiles("tpl.gohtml")
	err := tmpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
