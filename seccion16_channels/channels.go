package main

import (
	"fmt"
	"time"
)

func main() {

	var ch chan string = make(chan string)

	go func(ch chan string) {
		for i := 0; i <= 15; i++ { 
			ch <- fmt.Sprintf("mensaje: %v", i) // operador para enviar mensajes
		}
	}(ch)

	go func(ch chan string) {
		for { //ciclo for infinito
			mensaje := <-ch // operador para recivir mensajes
			fmt.Println(mensaje)
			time.Sleep(time.Second * 1) //se detiene la ejecución un segundo
		}
	}(ch)

	fmt.Println("\nOprime la tecla enter para deter el programa\n")

	var tecla string
	fmt.Scanln(&tecla)

}
