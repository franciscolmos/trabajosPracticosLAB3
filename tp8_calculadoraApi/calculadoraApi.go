package tp_8_calculadoraApi

import (
	"github.com/franciscolmos/trabajosPracticosLAB3/tp8_calculadoraApi/controller"
	"github.com/gin-gonic/gin"
)

func CalculadoraApi() {
	r := gin.Default()
	r.GET("/calc/sum", controller.Sumar)
	r.GET("/calc/res", controller.Restar)
	r.GET("/calc/mul", controller.Multiplicar)
	r.GET("/calc/div", controller.Dividir)
	r.Run(":8080")
}