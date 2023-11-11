package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	// you must have to be into the directory to run "go run main.go"
	t := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 100},
		{"Python", 50},
	})

	if err != nil {
		panic(err)
	}

}
