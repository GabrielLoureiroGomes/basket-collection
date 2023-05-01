package api

import "github.com/gin-gonic/gin"

func StartServer() {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Funcionando!",
		})
	})
	r.Run()
}