package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

type Hello struct {
	tmpl *template.Template
}

var hello Hello = Hello{}

type customer struct {
	ID      int
	Name    string
	Surname string
	Age     int
}

func datetimeFormat(t time.Time) string {
	return t.Format("2006/01/02")
}

func init() {
	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"fDate": datetimeFormat,
	}
	tmpl, err := template.New("func.html").Funcs(funcMap).ParseFiles("./templates/func.html")
	if err != nil {
		log.Fatalln(err)
	}
	hello.tmpl = tmpl
}

func main() {
	now := time.Now()
	err := hello.tmpl.Execute(os.Stdout, now)
	if err != nil {
		log.Fatalln(err)
	}

}
