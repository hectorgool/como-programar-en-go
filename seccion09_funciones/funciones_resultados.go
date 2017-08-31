/*
twitter@hector_gool
*/

package main

import ( 
	"fmt"
	"strings"
)

func main() {

	nombre, apellido := datos("Rodolfo", "Guzmán")
	fmt.Printf("Nombre: %v, Apellido: %v \n", nombre, apellido)

	// _ es un identificador en blaco, sirve para descartar valores
	n, _ := datos("Rodolfo", "Guzmán")

	// _ es un identificador en blaco, sirve para descartar valores
	_, a := datos("Rodolfo", "Guzmán")

	fmt.Printf("Nombre: %v \n", n)
	fmt.Printf("Apellido: %v \n", a)

}

func datos(nombre string, apellido string) (string, string){
	return strings.ToUpper(nombre), strings.ToUpper(apellido)
}