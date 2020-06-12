package tp_8_calculadoraApi

import (
"github.com/franciscolmos/trabajosPracticosLAB3/tp8_calculadoraApi/controller"
"github.com/gin-gonic/gin"
)

func CalculadoraApi() {
	r := gin.Default()
	r.GET("/sumar/:num1/:num2", controller.Sumar)
	r.GET("/restar/:num1/:num2", controller.Restar)
	r.GET("/multiplicar/:num1/:num2", controller.Multiplicar)
	r.GET("/dividir/:num1/:num2", controller.Dividir)
	r.Run(":3000")
}