package main

func main() {

	// Memoria -> Endereco -> Valor
	// variavel -> ponteiro que tem um endereco na memoria -> valor

	a := 10
	var ponteiro *int = &a
	*ponteiro = 20
	b := &a
	*b = 30
	println(a)
}
