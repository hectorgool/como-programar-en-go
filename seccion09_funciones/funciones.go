/*
twitter@hector_gool
*/

package main

import "fmt"

func main() {

	titulo()
	datos("Rodolfo", "santo@correo.com", 33)

}

func titulo(){
	fmt.Println("\n Esta es la función Título\n")
}

func datos(nombre string, correo string, edad int) {
	fmt.Printf("nombre: %s \n", nombre)
	fmt.Printf("correo: %s \n", correo)
	fmt.Printf("edad: %d \n", edad)
}