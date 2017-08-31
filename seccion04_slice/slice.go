/*
twitter@hector_gool
*/

package main

import "fmt"

func main() {

	var luchadores []string

	fmt.Println(luchadores)

	fmt.Println(len(luchadores))

	luchadores = append(luchadores, "Santo")
	luchadores = append(luchadores, "Blue Demon")
	luchadores = append(luchadores, "Mil Mascaras")

	fmt.Printf("Los luchadores son: %s \n", luchadores)

	fmt.Println(luchadores[2])

	for n, luchador := range luchadores{
		fmt.Printf(" %d %s \n", n, luchador)
	}

}