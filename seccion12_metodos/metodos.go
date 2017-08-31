package main

import "fmt"

type Parametros struct {
	A int
	B int
}

func main() {

	p := Parametros{
		A: 1,
		B: 2,
	}

	fmt.Printf("\n Los parametros son: %v y %v \n", p.A, p.B)

	fmt.Printf("\n El resultado de la suma es: %v \n", p.suma())

	fmt.Printf("\n El resultado de la multiplicación es: %v \n", p.multiplicacion())	

}

func (p Parametros) suma() int {
	return p.A + p.B
}

func (p Parametros) multiplicacion() int {
	return p.A * p.B
}
