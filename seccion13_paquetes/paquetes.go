package main

import (
	"fmt"
	"ejercicios/seccion12_paquetes/areas"
)

func main() {
	
	var radio, base, altura float64 = 10, 4, 5

	fmt.Printf("El radio es: %v \n", radio)
	fmt.Printf("La base es: %v \n", base)
	fmt.Printf("La altura es: %v \n", altura)

	fmt.Println()

	fmt.Printf("El área del rectángulo es = %v \n", areas.AreaRectangulo(base,altura))
	fmt.Printf("El área del triángulo es = %v \n", areas.AreaTriangulo(base,altura))
	fmt.Printf("El área del círculo es = %v \n", areas.AreaCirculo(radio))

}