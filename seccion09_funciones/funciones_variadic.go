/*
twitter@hector_gool
*/

package main

import ( 
	"fmt"
	"strings"
)

func main() {

	fmt.Printf("\n En Mayúsculas: %v \n", PasarAMayusculas("pera","uva"))
	fmt.Printf("\n En Mayúsculas: %v \n", PasarAMayusculas())
	fmt.Printf("\n En Mayúsculas: %v \n", PasarAMayusculas("manzana"))

	frutas := []string{"mango", "sandía", "platano"}

	fmt.Printf("\n En Mayúsculas: %v \n", PasarAMayusculas(frutas...))

}

func PasarAMayusculas(frutas ...string) []string{
	var resultado []string 
	for _, fruta := range frutas{
		resultado = append(resultado, strings.ToUpper(fruta))
	}
	return resultado
}