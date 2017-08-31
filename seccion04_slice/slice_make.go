/*
twitter@hector_gool
*/

package main

import "fmt"

func main() {

	s := make([]int, 3)

	fmt.Printf("el slice es: %d \n", s)

	s[0] = 1
	s[1] = 2
	s[2] = 3	

	fmt.Printf("el slice es: %d \n", s)

	s = append(s, 4,5,6,10)

	fmt.Printf("el slice es: %d \n", s)

	fmt.Printf("el tamaño del slice es: %d \n", len(s))

	for n, e := range s{
		fmt.Printf(" %d) \t %d \n", n, e)
	}

}