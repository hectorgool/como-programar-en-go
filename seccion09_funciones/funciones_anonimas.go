package main

import "fmt"

func main() {

	func(){
		fmt.Println("Esta es la función anónima")
	}()

	resultado := func( a, b int) int {
		return a * b
	}(3,4)

	fmt.Printf("\n El resultado es %v \n", resultado)

}
