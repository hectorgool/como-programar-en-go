package main

import (
	"fmt"
)

type Operaciones interface {
	suma() int
	multiplicacion() int
}

type Parametros struct {
	A int
	B int
}

func (p Parametros) suma() int {
	return p.A + p.B
}

func (p Parametros) multiplicacion() int {
	return p.A * p.B
}

func main() {

	var p Parametros

	p = Parametros{
		A: 1,
		B: 2,
	}

	fmt.Printf("\nLa suma de %v + %v es: %v\n", p.A, p.B, p.suma())
	fmt.Printf("\nLa multiplicación de %v + %v es: %v\n\n", p.A, p.B, p.multiplicacion())

}
