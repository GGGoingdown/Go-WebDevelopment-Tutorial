package main

import (
	"html/template"
	"log"
	"os"
)

type Hello struct {
	tmpl *template.Template
}

type customer struct {
	ID      int
	Name    string
	Surname string
	Age     int
}

func main() {
	hello := Hello{}
	tmpl, err := template.ParseFiles("./templates/hello.html")
	if err != nil {
		log.Fatalln(err)
	}
	hello.tmpl = tmpl
	user1 := customer{
		ID:      1,
		Name:    "eddie",
		Surname: "pai",
		Age:     30,
	}
	hello.tmpl.Execute(os.Stdout, user1)
}
