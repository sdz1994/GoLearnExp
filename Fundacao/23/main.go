package main

import "fmt"

func main() {
	//default
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//key value
	numeros := []string{"um", "dois", "tres"}
	for k, v := range numeros {
		println(k, v)
	}

	for _, v := range numeros {
		println(v)
	}

	for k, _ := range numeros {
		println(k)
	}

	//as while
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	//infinito
	for {
		println("Hello World")
	}

}

//Apenas FOR
