package main

import "fmt"

type Datos struct{
	nombre string
	email string
	edad int
	esInstructor bool
}

func main() {
	
	var d Datos

	fmt.Printf("Nombre: %s \n", d.nombre)
	fmt.Printf("Email: %s \n", d.email)
	fmt.Printf("edad: %d \n", d.edad)
	fmt.Printf("Es Instructor: %t \n", d.esInstructor)

	d.nombre = "Rodolfo"
	d.email = "Email"
	d.edad = 22
	d.esInstructor = true

	fmt.Println(d)

	d2 := Datos{ 
		nombre: "Emi", 
		email: "emi@correo.com", 
		edad: 12, 
		esInstructor: false,
	}

	fmt.Println(d2)

}