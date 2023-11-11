package main

import (
	"html/template"
	"net/http"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	// you must have to be into the directory to run "go run main.go"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))

		err := t.Execute(w, Cursos{
			{"Go", 40},
			{"Java", 100},
			{"Python", 50},
		})

		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8282", nil)

}
