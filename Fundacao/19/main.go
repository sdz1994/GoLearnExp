package main

import "fmt"

func main() {
	var minhaVar interface{} = "Wesley Willians"
	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res eh %v e o resultado de ok eh %v\n", res, ok)
	res2 := minhaVar.(int)
	fmt.Printf(" O valor de res2 eh %v\n", res2)
}
