package main

import (
	"fmt"
	"github.com/franciscolmos/trabajosPracticosLAB3/tp3_hello_word"
	"github.com/franciscolmos/trabajosPracticosLAB3/tp4_calculdora"
	"github.com/franciscolmos/trabajosPracticosLAB3/tp5_consumo_apis"
	"github.com/franciscolmos/trabajosPracticosLAB3/tp6_calculadora_hilos"
	"github.com/franciscolmos/trabajosPracticosLAB3/tp7_calculadora_test"
	"github.com/franciscolmos/trabajosPracticosLAB3/tp8_calculadoraApi"
)

func main(){
	var numeroTP int64
	fmt.Printf("Ingrese el numero del TP que quiere ejecutar: \n")
	fmt.Printf("3): Introduccion a Go \n")
	fmt.Printf("4): Variables y estructuras en Go \n")
	fmt.Printf("5): Consumo de APIs en Go \n")
	fmt.Printf("6): Concurrencia y Paralelirmos en Go \n")
	fmt.Printf("7): Testing \n")
	fmt.Printf("8): Creacion de una API \n\n ")
	fmt.Scan(&numeroTP)
	switch numeroTP {
	case 3:
		tp3_hello_word.HelloWord()
		break
	case 4:
		tp4_calculdora.Calculadora()
		break
	case 5:
		tp5_consumo_apis.ConsumoApi()
		break
	case 6:
		tp6_calculadora_hilos.CalculadoraHilos()
		break
	case 7:
		tp7_calculadora_test.CalculadoraTest()
		break
	case 8:
		tp_8_calculadoraApi.CalculadoraApi()
		break
	}
}
