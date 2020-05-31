package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// ParseFiles accepts name of template files to read and returns a pointer to template and error
	tpl, err := template.ParseFiles("templates/tpl.gohtml")
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}
}
