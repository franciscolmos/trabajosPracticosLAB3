package tp5_consumo_apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func ConsumoApi(){
	r := gin.Default()
	r.GET("https://api.mercadolibre.com/sites/MLA", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ "message": "" })
	})
	r.Run(":3000")
}

