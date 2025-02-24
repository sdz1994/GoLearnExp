package main

import (
	"html/template"
	"os"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"/home/sedevitiz/GolandProjects/goExpert/Pacotes-importantes/17/header.html",
		"/home/sedevitiz/GolandProjects/goExpert/Pacotes-importantes/17/content.html",
		"/home/sedevitiz/GolandProjects/goExpert/Pacotes-importantes/17/footer.html",
	}

	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})
	t = template.Must(t.ParseFiles(templates...))
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 60},
		{"Python", 80},
	})
	if err != nil {
		panic(err)
	}

}
