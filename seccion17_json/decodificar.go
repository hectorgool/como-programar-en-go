package main

import (
	"fmt"
	"encoding/json"
)

type Banda struct{
	Nombre string
	Vocalista string
	Bajista string
	Baterista string
	Guitarrista string
}

func main() {

	JSON :=`
		{
			"Nombre": "U2",
			"Vocalista": "Bono",
			"Bajista": "Adam Clayton",
			"Baterista": "Larry Mullen",
			"Guitarrista": "The Edge"
		}
	`

	var b Banda

	fmt.Printf("El objeto json es: %v \n", JSON)

	err := json.Unmarshal( []byte(JSON), &b )
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("La estructura en go es: %v \n", b)

}
