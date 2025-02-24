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
	t := template.Must(template.New("cursoTemplate)").Parse("Curso: {{.Nome}} - Carga Horaria: {{.CargaHoraria}} - 2"))
	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}

}
