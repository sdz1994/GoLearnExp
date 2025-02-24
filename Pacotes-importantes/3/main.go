package main

import "fmt"

func main() {
	//defer - atrasa pra ultima execucao, executa por ultimo
	fmt.Println("Primeira Linha")
	defer fmt.Println("Segunda Linha")
	fmt.Println("Terceira Linha")
}
