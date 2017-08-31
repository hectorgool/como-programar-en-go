package main

import "fmt"

func main() {

	m := map[string]string{
		"nombre" : "Rodolfo",
		"apellido" : "Guzmán",
	}

	fmt.Printf("El mapa es: %s \n", m)

	fmt.Printf("\n El nombre es: %s \n", m["nombre"])

	m["nickname"] = "Santo"

	fmt.Printf("\n El tipo es: %T \n\n", m)

	for k, v := range m{
		fmt.Printf("%s \t %s \n", k, v)
	}

}