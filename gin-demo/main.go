package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("v1")

	v1.GET("/msg", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})

	})

	v1.POST("/add", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "add ok",
		})

	})

	r.Run(":8081")
}
