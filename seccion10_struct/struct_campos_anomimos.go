package main

import "fmt"

type Canciones struct {
	Nombres []string
}

type Banda struct {
	Nombre    string
	Canciones // campo sin nombre de tipo Canciones, "campo anónimo"
}

func main() {
	
	banda1 := Banda{
		Nombre: "U2",
		Canciones: Canciones{
			Nombres: []string{"One", "Vertigo", "Elevation", "Stay"},
		},		
	}

	fmt.Printf("\n La estructura completa es: %v \n", banda1)
	fmt.Printf("\n El nombre de la banda es: %v \n", banda1.Nombre)
	fmt.Printf("\n Las canciones son: %v \n", banda1.Canciones.Nombres)
	for i, miembro := range banda1.Canciones.Nombres {
		i++
		fmt.Printf(" %v) %v \n", i, miembro)
	}

	fmt.Println("----------------------------------------------------------------")

	banda2 := Banda{
		"Beatles",
		Canciones{
			[]string{"Help", "Yesterday", "Revolution", "Blackbird"},
		},
	}

	fmt.Printf("\n La estructura completa es: %v \n", banda2)
	fmt.Printf("\n El nombre de la banda es: %v \n", banda2.Nombre)
	fmt.Printf("\n Las canciones son: %v \n", banda2.Canciones.Nombres)
	for i, miembro := range banda2.Canciones.Nombres {
		i++
		fmt.Printf(" %v) %v \n", i, miembro)
	}	
}