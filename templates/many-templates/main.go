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
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	// you need to put the main file that it will receive the arguments, like is showed in t.Execute(
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("content.html").ParseFiles(templates...))

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
