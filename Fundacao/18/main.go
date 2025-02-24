package main

import "fmt"

func main() {

	var x interface{} = 10
	var y interface{} = "Hello, World!"

	showType(x)
	showType(y)

}

func showType(t interface{}) {
	fmt.Printf("O tipo de variavel eh %T e o valor eh %v\n", t, t)
}
