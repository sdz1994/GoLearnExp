package main

import (
	"curso-go/matematica"
	"fmt"
)

func main() {
	s := matematica.Soma(10, 30)
	carro := matematica.Carro{Marca: "fiat"}
	fmt.Println("O resultado: ", s)
	fmt.Println(carro.Andar())
	fmt.Println(carro.Marca)
	fmt.Println(matematica.A)
}

//Letra Maiuscula em func/var/struct para acessos de pacotes externos, minusculo interno
//Letra Maiuscula == public
//Letra Minuscula == private
