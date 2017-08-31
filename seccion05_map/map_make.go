/*
twitter@hector_gool
*/

package main

import "fmt"

func main() {

	m := make(map[string]string)

	fmt.Printf("el mapa es: %v \n", m)

	m["nombre"] = "Rodolfo"
	m["apellido"] = "Guzmán"
	m["nickname"] = "Santo"

	fmt.Printf("el mapa es: %s \n", m)

	fmt.Printf("el nombre es: %s \n", m["nombre"])

	if name, valida := m["nickname"]; valida{
		fmt.Println(name, valida)
	}

	fmt.Printf("\n El tipo es: %T \n", m)

	contador := 0
	for k, v := range m{
		contador ++
		fmt.Printf("%d \t %s \t %s \n", contador, k, v)
	}

}