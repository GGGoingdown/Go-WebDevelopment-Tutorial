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
	ID    int
	Fname string
	Lname string
	Age   int
}

func main() {
	hello := Hello{}
	tmpl, err := template.ParseFiles("./templates/loop.html")
	if err != nil {
		log.Fatalln(err)
	}
	hello.tmpl = tmpl
	user1 := customer{
		ID:    1,
		Fname: "eddie",
		Lname: "pai",
		Age:   30,
	}
	user2 := customer{
		ID:    1,
		Fname: "peter",
		Lname: "pan",
		Age:   30,
	}

	users := []customer{user1, user2}
	hello.tmpl.Execute(os.Stdout, users)
}
