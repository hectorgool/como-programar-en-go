/*
twitter@hector_gool
*/

package main

import "fmt"

func main() {
	
	var frutas [3]string
	fmt.Println(frutas)

	frutas[0] = "Pera"
	frutas[1] = "Manzana"
	frutas[2] = "Uva"

	fmt.Println(frutas)

	fmt.Println(frutas[2])

	for n, fruta := range frutas{
		fmt.Printf("%v - fruta %v \n", n, fruta)
	}

}