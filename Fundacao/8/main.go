package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sumB(5, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}

func sum(a int, b int) int {
	return a + b
}

func sumB(a int, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma eh maior que 50")
	}
	return sum(a, b), nil
}
