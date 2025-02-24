package main

import "fmt"

type Cliente struct {
	nome string
	Conta
}

type Conta struct {
	saldo int
}

func (c Cliente) andou() {
	c.nome = "Wesley Willians"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

func newConta() *Conta {
	return &Conta{saldo: 0}
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func main() {
	//conta := Conta{saldo: 100}
	//conta.simular(200)
	conta1 := newConta()
	conta2 := newConta()
	conta1.simular(200)
	println(conta1.saldo)
	println(conta2.saldo)
}
