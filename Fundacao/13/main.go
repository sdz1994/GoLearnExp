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

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

func main() {
	wesley := Cliente{Nome: "Wesley", Idade: 30, Ativo: true}
	wesley.Desativar()

	fmt.Printf("O status de %s eh %s", wesley.Nome, wesley.Ativo)

}
