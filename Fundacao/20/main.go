package main

func SomaInteiro(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

type MyNumber int

// constraint
type Number interface {
	~int | ~float64
}

// com generics
func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Wesley": 1000, "Joao": 2000, "Maria": 3000}
	m2 := map[string]float64{"Wesley": 100.10, "Joao": 200.20, "Maria": 300.30}
	m3 := map[string]MyNumber{"Wesley": 1000, "Joao": 2000, "Maria": 3000}
	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10.))
}
