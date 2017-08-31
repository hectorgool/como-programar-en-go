package main

import (
	"fmt"
	"encoding/json"
)

type Banda struct{
	Nombre string
	Vocalista string
	Bajista string
	Baterista string
	Guitarrista string
}

func main() {

	banda := Banda{ 
		Nombre: "U2",
		Vocalista: "Bono",
		Bajista: "Adam Clayton",
		Baterista: "Larry Mullen",
		Guitarrista: "The Edge",
	}

	fmt.Printf("La estructura en go es: %v \n", banda)

	datos, _ := json.Marshal(banda)

	fmt.Println("La estructura en go codificada a un objeto json es: ")
	fmt.Println(string(datos))

}
