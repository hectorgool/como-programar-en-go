/*
twitter@hector_gool
*/

package main

import "fmt"

func main() {

	fmt.Println("\nIngresa un número del 1 al 7, para decirte a que dia corresponde\n")	

	var dia int

	fmt.Scan(&dia)

	switch dia {
		case 1: fmt.Println("Lunes")
		case 2: fmt.Println("Mártes")
		case 3: fmt.Println("Miércoles")
		case 4: fmt.Println("Jueves")
		case 5: fmt.Println("Viernes")
		case 6: fmt.Println("Sábado")
		case 7: fmt.Println("Domingo")
	}

	/*
	if dia == 1 {
		fmt.Println("Lunes")
	} else if dia == 2 {
		fmt.Println("Mártes")
	} else if dia == 3 {
		fmt.Println("Miércoles")
	} else if dia == 4 {
		fmt.Println("Jueves")
	} else if dia == 5 {
		fmt.Println("Viernes")
	} else if dia == 6 {
		fmt.Println("Sábado")
	} else if dia == 7 {
		fmt.Println("Domingo")
	} else {
		fmt.Println("El dato ingresado no es válido")
	}
	*/

}
