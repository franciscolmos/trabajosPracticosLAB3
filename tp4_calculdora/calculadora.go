package tp4_calculdora

import (
	"fmt"
	"github.com/franciscolmos/trabajosPracticosLAB3/tp4_calculdora/operaciones"
)

func Calculadora()  {
	var num1, num2 float64
	var operacion string
	var operandos []float64
	var respuesta string
	var bug bool
	var offset int
	fmt.Printf("Ingrese la operación a realizar: \n")
	fmt.Printf("sumar \n")
	fmt.Printf("restar \n")
	fmt.Printf("multiplicar \n")
	fmt.Printf("dividir \n")
	fmt.Scan(&operacion)
	fmt.Printf("Ingrese dos operandos: \n")
	fmt.Scan(&num1, &num2)
	operandos = []float64{num1,num2}
	fmt.Printf("Desea buguear la calculadora?: (S/N)")
	fmt.Scan(&respuesta)
	if(respuesta == "S"){
		bug = true
		fmt.Printf("Ingrese OFFSET del Bug:")
		fmt.Scan(&offset)
	}
	switch operacion {
	case "sumar":
		fmt.Print(operaciones.Sumar(bug,offset,operandos...))
	case "restar":
		fmt.Print(operaciones.Restar(bug,offset,operandos...))
	case "multiplicar":
		fmt.Print(operaciones.Multiplicar(bug,offset,operandos...))
	case "dividir":
		fmt.Print(operaciones.Division(bug,offset,operandos...))
	default:
		fmt.Print("Operracion inválida.")
	}
}
