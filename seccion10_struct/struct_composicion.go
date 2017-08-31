package main

import (
	"fmt"
)

type Integrante struct {
	Nombre string
}

type Banda struct {
	Nombre      string
	Integrantes []Integrante 
}

func main() {

	banda1 := Banda{
		Nombre: "U2",
		Integrantes: []Integrante{
			Integrante{
				Nombre: "Bono",
			},
			Integrante{
				Nombre: "The Edge",
			},
			Integrante{
				Nombre: "Adam Clayton",
			},
			Integrante{
				Nombre: "Larry Mullen",
			},
		},
	}

	banda2 := Banda{
		"Beatles",
		[]Integrante{
			Integrante{
				"John Lennon",
			},
			Integrante{
				"George Harrison",
			},
			Integrante{
				"Ringo Starr",
			},
			Integrante{
				"Paul McCartney",
			},
		},
	}

	fmt.Printf("\nLa estructura es: %v\n", banda1)
	fmt.Printf("La Banda es: %v \n", banda1.Nombre)
	fmt.Printf("Los Integrantes son: %v \n", banda1.Integrantes)
	fmt.Printf("Un integrantes es: %v \n", banda1.Integrantes[0])
	fmt.Printf("Un integrantes es: %v \n", banda1.Integrantes[0].Nombre)

	fmt.Println("-----------------------------------------")

	fmt.Printf("\nLa estructura es: %v\n", banda2)
	fmt.Printf("La Banda es: %v \n", banda2.Nombre)
	fmt.Printf("Los Integrantes son: %v \n", banda2.Integrantes)
	fmt.Printf("Un integrantes es: %v \n", banda2.Integrantes[0])
	fmt.Printf("Un integrantes es: %v \n", banda2.Integrantes[0].Nombre)

}
