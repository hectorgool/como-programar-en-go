/*
twitter@hector_gool
*/

package main

import (
	"fmt"
)

func main() {

	var nombre string
	var edad int
	var version float32
	var aprobado bool

	fmt.Printf("%q \n", nombre)
	fmt.Printf("%d \n", edad)
	fmt.Printf("%f \n", version)
	fmt.Printf("%t \n", aprobado)

	fmt.Println("\n Asignación de valores a variables: \n")

	nombre = "Santo"
	fmt.Printf("%v \n", nombre)

	edad = 33
	fmt.Printf("%v \n", edad)

	version = 1.1
	fmt.Printf("%v \n", version)

	aprobado = true
	fmt.Printf("%v \n", aprobado)

}