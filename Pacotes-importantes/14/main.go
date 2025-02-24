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
	t := template.Must(template.New("content.html").ParseFiles("/home/sedevitiz/GolandProjects/goExpert/Pacotes-importantes/14/content.html"))
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 60},
		{"Python", 80},
	})
	if err != nil {
		panic(err)
	}
}
