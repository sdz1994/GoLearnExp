package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome    string
	Idade   int
	Ativo   bool
	Address Endereco
}

func main() {
	wesley := Cliente{Nome: "Wesley", Idade: 30, Ativo: true}
	wesley.Ativo = false
	wesley.Address.Cidade = " Sao Paulo"

	fmt.Printf(wesley.Nome)

}
