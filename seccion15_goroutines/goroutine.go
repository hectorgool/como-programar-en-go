package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

const n int = 100

func main(){

	wg.Add(2)

	for i := 1; i <= 10; i++ {
		go tabla_multiplicar(i)
	}
	//go tabla_multiplicar(1)
	//go tabla_multiplicar(2)

	wg.Wait()

	fmt.Println("Fin")

}

func tabla_multiplicar( parametro int ){
	fmt.Printf("La tabla del %d es: \n", parametro)
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d x %d = %d \n", parametro, i, parametro * i )
	}
	fmt.Println()
	defer wg.Done()

}