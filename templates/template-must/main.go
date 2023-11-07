package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"Go", 40}

	// TEMPLATE MUST, that function is more clear than conventional
	templateMust := template.Must(template.New("Curso Template").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}"))
	err := templateMust.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
