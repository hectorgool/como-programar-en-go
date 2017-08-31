package main

import "fmt"

func main() {
	
	var colores = [...]string{"rojo", "amarillo", "azul", "rosa"}

	fmt.Println(colores)

	fmt.Println(len(colores))

	colores[1] = "verde"

	fmt.Println(colores)

}