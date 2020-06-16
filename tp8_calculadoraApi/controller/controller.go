package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Sumar(c *gin.Context) {
	a := c.Query("a")
	b := c.Query("b")
	if a == "" || b == "" {
		c.String(400, "Missing params a,b")
		return
	}
	num1, error1 := strconv.ParseInt(a, 10, 64)
	num2, error2 := strconv.ParseInt(b, 10, 64)
	if error1 != nil || error2 != nil {
		c.String(400, "Invalid params a,b")
		return
	}
	result := num1 + num2
	c.JSON(http.StatusOK, gin.H{"Result": result})
}

func Restar(c *gin.Context) {
	a := c.Query("a")
	b := c.Query("b")
	if a == "" ||  b == "" {
		c.String(400, "Missing params a,b")
		return
	}
	num1, error1 := strconv.ParseInt(a, 10, 64)
	num2, error2 := strconv.ParseInt(b, 10, 64)
	if error1 != nil || error2 != nil {
		c.String(400, "Invalid params a,b")
		return
	}
	result := num1 - num2
	c.JSON(http.StatusOK, gin.H{"Result": result})
}


func Multiplicar(c *gin.Context) {
	a := c.Query("a")
	b := c.Query("b")
	if a == "" || b == "" {
		c.String(400, "Missing params a,b")
		return
	}
	num1, error1 := strconv.ParseInt(a, 10, 64)
	num2, error2 := strconv.ParseInt(b, 10, 64)
	if error1 != nil || error2 != nil {
		c.String(400, "Invalid params a,b")
		return
	}
	result := num1 * num2
	c.JSON(http.StatusOK, gin.H{"Result": result})
}


func Dividir(c *gin.Context) {
	a := c.Query("a")
	b := c.Query("b")
	if a == "" || b == "" {
		c.String(400, "Missing params a,b")
		return
	}
	num1, error1 := strconv.ParseInt(a, 10, 64)
	num2, error2 := strconv.ParseInt(b, 10, 64)
	if error1 != nil || error2 != nil {
		c.String(400, "Invalid params a,b")
		return
	}
	if num2 == 0 {
		c.String(400, "Invalid param b, cannot be 0 in div")
		return
	}
	result := num1 / num2
	c.JSON(http.StatusOK, gin.H{"Result": result})
}